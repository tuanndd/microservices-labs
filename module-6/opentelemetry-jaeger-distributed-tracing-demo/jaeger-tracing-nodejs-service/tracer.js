const { LogLevel } = require("@opentelemetry/core");
const { NodeTracerProvider } = require("@opentelemetry/node");
const { SimpleSpanProcessor } = require("@opentelemetry/tracing");
const { JaegerExporter } = require('@opentelemetry/exporter-jaeger');

const serviceName = process.env.SERVICE_NAME || "nodejs-service";
const jaegerAgentHost = process.env.JAEGER_AGENT_HOST || 'localhost';
const jaegerAgentPort = process.env.JAEGER_AGENT_PORT || 'localhost';


const provider = new NodeTracerProvider({
    plugins: {
        http: {
        enabled: true,
        path: '@opentelemetry/plugin-http',
            ignoreIncomingPaths: [
            '/',
            '/health'
            ]
        },
    },
    logLevel: LogLevel.ERROR
});

provider.register();

provider.addSpanProcessor(
  new SimpleSpanProcessor(
    new JaegerExporter({
        serviceName: serviceName,
        host: jaegerAgentHost,
        port: jaegerAgentPort
    })
  )
);

console.log("Tracing initialized");
