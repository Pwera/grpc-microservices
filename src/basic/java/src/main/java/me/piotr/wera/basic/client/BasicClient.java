package me.piotr.wera.basic.client;

import com.google.common.collect.Streams;
import com.google.common.util.concurrent.Uninterruptibles;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;
import io.grpc.stub.StreamObserver;
import io.opencensus.trace.Status;
import me.piotr.wera.basic.BasicRequest;
import me.piotr.wera.basic.BasicResponse;
import me.piotr.wera.basic.BasicServiceGrpc;
import me.piotr.wera.basic.GrpcCommons;

import java.util.concurrent.TimeUnit;
import java.util.stream.Stream;

public class BasicClient implements GrpcCommons {

    private static int messageID = 0;

    public static void main(String[] args) {
        new BasicClient().go();
    }

    public void go() {
        System.out.println("Starting gRPC Client");
        final ManagedChannel channel = createChannel();

        final BasicServiceGrpc.BasicServiceBlockingStub syncClient = BasicServiceGrpc.newBlockingStub(channel);

        System.out.println("Calling gRPC Unary Method");
        try {
            Stream.of(
                    syncClient.withDeadline(deadline)
                            .myServiceUnary(constructRequest())).forEach(this::responseFromServer);
        } catch (StatusRuntimeException e) {
            if (Status.INVALID_ARGUMENT.equals(e.getStatus())) {
                System.out.println("INVALID_ARGUMENT error");
            } else if (Status.DEADLINE_EXCEEDED.equals(e.getStatus())) {
                System.out.println("DEADLINE_EXCEEDED error");
            } else {
                e.printStackTrace();
            }
        }


        System.out.println("Calling gRPC Server Streaming Method");
        Streams.stream(syncClient.myServiceServerStreaming(constructRequest())).forEach(this::responseFromServer);

        Uninterruptibles.sleepUninterruptibly(1000L, TimeUnit.MILLISECONDS);
        System.out.println("Calling gRPC Client Streaming Method");
        final BasicServiceGrpc.BasicServiceStub customerServiceStub = BasicServiceGrpc.newStub(channel);
        final StreamObserver<BasicRequest> BasicRequestStreamObserver = customerServiceStub.myServiceClientStreaming(new BasicResponseStreamObserver());
        BasicRequestStreamObserver.onNext(constructRequest());
        BasicRequestStreamObserver.onNext(constructRequest());
        BasicRequestStreamObserver.onNext(constructRequest());
        BasicRequestStreamObserver.onCompleted();

        Uninterruptibles.sleepUninterruptibly(1000L, TimeUnit.MILLISECONDS);
        System.out.println("Calling gRPC BiDi Streaming Method");
        final StreamObserver<BasicRequest> BasicRequestBiDiObserver = customerServiceStub.myServiceClientStreaming(new BasicResponseStreamObserver());
        BasicRequestBiDiObserver.onNext(constructRequest());
        BasicRequestBiDiObserver.onNext(constructRequest());
        BasicRequestBiDiObserver.onNext(constructRequest());
        BasicRequestBiDiObserver.onCompleted();


        System.out.println("Shutting down channel");
        Uninterruptibles.sleepUninterruptibly(4000L, TimeUnit.MILLISECONDS);
        channel.shutdown();

    }

    private BasicRequest constructRequest() {
        return BasicRequest
                .newBuilder()
                .setValue(messageID % 2 == 0 ? "expected" : "not-exxpected")
                .setId(1)
                .build();
    }

    private void responseFromServer(BasicResponse BasicResponse) {
        System.out.println("Response from server: " + BasicResponse.getStatus() + " " + BasicResponse.getId());

    }

    private class BasicResponseStreamObserver implements StreamObserver<BasicResponse> {
        @Override
        public void onNext(BasicResponse value) {
            Stream.of("Received message from server ", value).forEach(System.out::println);
        }

        @Override
        public void onError(Throwable t) {

        }

        @Override
        public void onCompleted() {

        }
    }

    private ManagedChannel createChannel() {
        return ManagedChannelBuilder.forAddress(GRPC_SERVER_HOST, GRPC_SERVER_PORT)
                .usePlaintext()
                .build();
    }


}
