FROM openjdk:11.0.7-jre-slim

ARG JAR_FILE=build/libs/jaeger-tracing-java-service-0.0.1-SNAPSHOT.jar

COPY opentelemetry-javaagent-all.jar opentelemetry-javaagent-all.jar
COPY ${JAR_FILE} app.jar

ENTRYPOINT ["java", "-javaagent:opentelemetry-javaagent-all.jar", "-jar","/app.jar"]