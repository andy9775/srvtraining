## Stage 6
* Added tracing package with support from opencensus.io
* Generate trace id on request or find existing in request.
* Handlers add a span to the trace.
* Updated main to start the tracing.