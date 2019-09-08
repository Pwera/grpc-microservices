package me.piotr.wera.common;

import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.protobuf.services.ProtoReflectionService;
import me.piotr.wera.blog.server.BlogServer;

import java.io.File;

public interface GrpcServerCommon {

    int GRPC_SERVER_PORT = 50051;
    boolean USE_SSL = true;
    File CERT_CHAIN = new File("ssl/server.crt");
    File PRIVATE_KEY = new File("ssl/server.pem");


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
}
