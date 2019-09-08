package me.piotr.wera.common;

import io.grpc.*;
import io.grpc.netty.shaded.io.grpc.netty.GrpcSslContexts;
import io.grpc.netty.shaded.io.grpc.netty.NettyChannelBuilder;
import io.grpc.protobuf.services.ProtoReflectionService;
import io.vavr.control.Try;
import me.piotr.wera.blog.server.BlogServer;

import java.io.File;
import java.util.concurrent.TimeUnit;

public interface GrpcCommons {

    int GRPC_SERVER_PORT = 50051;
    String GRPC_SERVER_HOST = "localhost";
    boolean USE_SSL = true;
    File CERT_CHAIN = new File("ssl/server.crt");
    File PRIVATE_KEY = new File("ssl/server.pem");
    File SSL_FILE = new File("ssl/ca.crt");
    Deadline deadline = Deadline.after(5000, TimeUnit.MILLISECONDS);


    default Server createServer(boolean useSSL) {
        final ServerBuilder builder = ServerBuilder.forPort(GRPC_SERVER_PORT)
                .addService(new BlogServer())
                .addService(ProtoReflectionService.newInstance());

        if (useSSL) {
            builder.useTransportSecurity(CERT_CHAIN, PRIVATE_KEY);

        }
        return builder.build();
    }

    default void onServerShutdown(Server server) {
        System.out.println("onServerShutdown triggered");
        server.shutdown();
        System.out.println("onServerShutdown executed");
    }
    default ManagedChannel createChannel(boolean useSSL) {
        if (useSSL) {
            return NettyChannelBuilder.forAddress(GRPC_SERVER_HOST, GRPC_SERVER_PORT)
                    .sslContext(Try.of(() ->
                            GrpcSslContexts.forClient().trustManager(SSL_FILE))
                            .mapTry(sslContextBuilder -> sslContextBuilder.build()).get())
                    .build();
        } else {
            return ManagedChannelBuilder.forAddress(GRPC_SERVER_HOST, GRPC_SERVER_PORT)
                    .usePlaintext()
                    .build();
        }
    }
}
