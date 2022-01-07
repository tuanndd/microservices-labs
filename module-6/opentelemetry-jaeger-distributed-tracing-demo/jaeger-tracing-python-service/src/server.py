from flask import Flask
import json
import os
import sys
import mysql.connector

from opentelemetry.instrumentation.mysql import MySQLInstrumentor
from opentelemetry import trace
from opentelemetry.exporter.jaeger.thrift import JaegerExporter
from opentelemetry.sdk.resources import SERVICE_NAME, Resource
from opentelemetry.instrumentation.flask import FlaskInstrumentor
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor

app = Flask(__name__)
trace.set_tracer_provider(
    TracerProvider(
        resource=Resource.create({SERVICE_NAME: os.getenv('OTEL_JAEGER_SERVICE_NAME', "python-service")})
    )
)
tracer = trace.get_tracer(__name__)

# SpanExporter receives the spans and send them to the target location.
exporter = JaegerExporter(
    agent_host_name=os.getenv('OTEL_JAEGER_AGENT_HOST', "localhost"),
    agent_port=6831
)

span_processor = BatchSpanProcessor(exporter)
trace.get_tracer_provider().add_span_processor(span_processor)

FlaskInstrumentor().instrument_app(app)
MySQLInstrumentor().instrument()

# configuration used to connect to MariaDB
config = {
    'host': os.getenv('MYSQL_HOST', "localhost"),
    'port': 3306,
    'user': os.getenv('MYSQL_USER', "admin"),
    'password': os.getenv('MYSQL_PASSWORD', "password"),
    'database': 'salary_amount'
}

@app.route("/salary-amount-for-grade/<grade>")
def getSalaryGrade(grade):
    print("request recieved")
    with tracer.start_as_current_span('salary-amount-for-grade'):
        try:
            conn = mysql.connector.connect(**config)
            cursor = conn.cursor()
            cursor.execute("select minimum, maximum from salary_amount where grade = '" + grade + "'")

            row_headers=[x[0] for x in cursor.description]
            rv = cursor.fetchall()
            json_data=[]
            for result in rv:
                print(result)
                json_data.append(dict(zip(row_headers,result)))
            return json.dumps(json_data[0])
        except mysql.connector.Error as e:
            print(f"Error connecting to MySQL Platform: {e}")
            sys.exit(1)

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8092)
