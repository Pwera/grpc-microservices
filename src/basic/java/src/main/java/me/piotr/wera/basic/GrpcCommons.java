package me.piotr.wera.basic;

import io.grpc.Deadline;

import java.util.concurrent.TimeUnit;

public interface GrpcCommons {

    int GRPC_SERVER_PORT = 50051;
    String GRPC_SERVER_HOST = "localhost";
    Deadline deadline = Deadline.after(5000, TimeUnit.MILLISECONDS);

}
