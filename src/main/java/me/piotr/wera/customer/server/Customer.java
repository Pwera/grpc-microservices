package me.piotr.wera.customer.server;

import io.grpc.Context;
import io.grpc.Server;
import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import me.piotr.wera.common.GrpcCommons;
import me.piotr.wera.customer.CustomerRequest;
import me.piotr.wera.customer.CustomerResponse;
import me.piotr.wera.customer.CustomerServiceGrpc;

import java.util.stream.Stream;

public class Customer extends CustomerServiceGrpc.CustomerServiceImplBase implements GrpcCommons {

    private static final String EXPECTED_VALUE = "expected";
    private static int messageID = 0;

    public static void main(String[] args) throws Exception {
        new Customer().go();
    }

    public void go() throws Exception {
        final Server server = createServer(USE_SSL);
        server.start();
        System.out.println("Customer listening for connection on port " + GRPC_SERVER_PORT);
        Runtime.getRuntime().addShutdownHook(new Thread(() -> onServerShutdown(server)));
        server.awaitTermination();
    }

    @Override
    public void myServiceUnary(CustomerRequest request, StreamObserver<CustomerResponse> responseObserver) {
        Stream.of(request, this, Context.current().toString(), "myService function").forEach(System.out::println);

        if (request.getId() < 0) {
            responseObserver.onError(Status.INVALID_ARGUMENT
                    .withDescription("Id was lower then 0")
                    .augmentDescription("ID: " + request.getId())
                    .asRuntimeException());
        }
        final CustomerResponse res = constructCustomerResponse(EXPECTED_VALUE.equals(request.getValue()));
        responseObserver.onNext(res);
        responseObserver.onCompleted();
    }

    @Override
    public void myServiceServerStreaming(CustomerRequest request, StreamObserver<CustomerResponse> responseObserver) {
        Stream.of(request, this, Context.current().toString(), "myServiceServerStreaming function").forEach(System.out::println);

        Stream.of(true, false, true, true)
                .map(s -> constructCustomerResponse(s))
                .forEach(responseObserver::onNext);


        responseObserver.onCompleted();
    }

    private CustomerResponse constructCustomerResponse(boolean status) {
        return CustomerResponse.newBuilder()
                .setStatus(status)
                .setId(messageID++)
                .build();
    }

    @Override
    public StreamObserver<CustomerRequest> myServiceClientStreaming(StreamObserver<CustomerResponse> responseObserver) {
        Stream.of(responseObserver, this, Context.current(), "myServiceClientStreaming function").forEach(System.out::println);
        return new CustomerRequestStreamObserver(responseObserver);
    }

    private class CustomerRequestStreamObserver implements StreamObserver<CustomerRequest> {

        final StreamObserver<CustomerResponse> responseObserver;
        String result = "";
        CustomerRequest latestCustomerRequest;

        public CustomerRequestStreamObserver(StreamObserver<CustomerResponse> responseObserver) {
            this.responseObserver = responseObserver;
        }

        @Override
        public void onNext(CustomerRequest req) {
            result += "Hello " + req.getValue() + "\n";
            latestCustomerRequest = req;
        }

        @Override
        public void onError(Throwable t) {

        }

        @Override
        public void onCompleted() {
            responseObserver.onNext(constructCustomerResponse(false));
        }
    }

    @Override
    public StreamObserver<CustomerRequest> myServiceBiDiStreaming(StreamObserver<CustomerResponse> responseObserver) {
        return new CustomerRequestStreamObserver2(responseObserver);
    }

    private class CustomerRequestStreamObserver2 implements StreamObserver<CustomerRequest> {

        final StreamObserver<CustomerResponse> responseObserver;

        public CustomerRequestStreamObserver2(StreamObserver<CustomerResponse> responseObserver) {
            this.responseObserver = responseObserver;
        }

        @Override
        public void onNext(CustomerRequest req) {
            responseObserver.onNext(constructCustomerResponse(req.getId() % 2 == 0));
        }

        @Override
        public void onError(Throwable t) {

        }

        @Override
        public void onCompleted() {
            responseObserver.onCompleted();
        }
    }

}
