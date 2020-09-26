package me.piotr.wera.basic.server;

import io.grpc.Context;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.Status;
import io.grpc.protobuf.services.ProtoReflectionService;
import io.grpc.stub.StreamObserver;
import me.piotr.wera.basic.BasicRequest;
import me.piotr.wera.basic.BasicResponse;
import me.piotr.wera.basic.BasicServiceGrpc;

import java.util.stream.Stream;

public class BasicServer extends BasicServiceGrpc.BasicServiceImplBase {

    private static final String EXPECTED_VALUE = "expected";
    private static final int GRPC_SERVER_PORT = 50051;
    private static int messageID = 0;

    public static void main(String[] args) throws Exception {
        new BasicServer().go();
    }

    public void go() throws Exception {
        final Server server = ServerBuilder.forPort(GRPC_SERVER_PORT)
                .addService(new BasicServer())
                .addService(ProtoReflectionService.newInstance()).build();
        server.start();
        System.out.println("Server listening on port " + GRPC_SERVER_PORT);
        Runtime.getRuntime().addShutdownHook(new Thread(() -> onServerShutdown(server)));
        server.awaitTermination();
    }

    @Override
    public void myServiceUnary(BasicRequest request, StreamObserver<BasicResponse> responseObserver) {
        System.out.printf("Got a request from {} through Unary request", Context.current().toString());

        responseObserver.onNext(constructBasicResponse("Unary BasicResponse"));
        responseObserver.onCompleted();
    }

    @Override
    public void myServiceServerStreaming(BasicRequest request, StreamObserver<BasicResponse> responseObserver) {
        System.out.printf("Got a request from {} through Server Streaming request", Context.current().toString());

        Stream.of("!", "@", "#", "$")
                .map(this::constructBasicResponse)
                .forEach(responseObserver::onNext);

        responseObserver.onCompleted();
    }

    @Override
    public StreamObserver<BasicRequest> myServiceClientStreaming(StreamObserver<BasicResponse> responseObserver) {
        System.out.printf("Got a request from {} through Client Streaming request", Context.current().toString());
        return new BasicRequestStreamObserver(responseObserver);
    }

    @Override
    public StreamObserver<BasicRequest> myServiceBiDiStreaming(StreamObserver<BasicResponse> responseObserver) {
        System.out.printf("Got a request from {} through Bidirectional Streaming request", Context.current().toString());
        return new BasicRequestStreamObserver2(responseObserver);
    }


    private BasicResponse constructBasicResponse(String status) {
        return BasicResponse.newBuilder()
                .setStatus(status)
                .setId(messageID++)
                .build();
    }


    private class BasicRequestStreamObserver implements StreamObserver<BasicRequest> {

        final StreamObserver<BasicResponse> responseObserver;
        String result = "";
        BasicRequest latestBasicRequest;

        public BasicRequestStreamObserver(StreamObserver<BasicResponse> responseObserver) {
            this.responseObserver = responseObserver;
        }

        @Override
        public void onNext(BasicRequest req) {
            result += "Received " + req.getValue() + "\n";
            latestBasicRequest = req;
        }

        @Override
        public void onError(Throwable t) {

        }

        @Override
        public void onCompleted() {
            responseObserver.onNext(constructBasicResponse("Client streaming BasicResponse"));
        }
    }


    private class BasicRequestStreamObserver2 implements StreamObserver<BasicRequest> {

        private final StreamObserver<BasicResponse> responseObserver;

        public BasicRequestStreamObserver2(StreamObserver<BasicResponse> responseObserver) {
            this.responseObserver = responseObserver;
        }

        @Override
        public void onNext(BasicRequest req) {
            responseObserver.onNext(constructBasicResponse(EXPECTED_VALUE));
        }

        @Override
        public void onError(Throwable t) {

        }

        @Override
        public void onCompleted() {
            responseObserver.onCompleted();
        }
    }

    void onServerShutdown(Server server) {
        System.out.println("onServerShutdown triggered");
        server.shutdown();
        System.out.println("onServerShutdown executed");
    }

}
