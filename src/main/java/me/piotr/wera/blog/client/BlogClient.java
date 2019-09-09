package me.piotr.wera.blog.client;

import com.google.common.base.Stopwatch;
import com.google.common.collect.Streams;
import com.google.common.util.concurrent.Uninterruptibles;
import io.grpc.ManagedChannel;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import me.piotr.wera.*;
import me.piotr.wera.common.GrpcCommons;

import java.util.concurrent.TimeUnit;
import java.util.stream.Stream;

public class BlogClient implements GrpcCommons {

    public static void main(String[] args) {
        new BlogClient().go();
    }

    public void go() {
        System.out.println("Starting gRPC Client");
        final ManagedChannel channel = createChannel(USE_SSL);

        final BlogServiceGrpc.BlogServiceBlockingStub syncClient = BlogServiceGrpc.newBlockingStub(channel);

        System.out.println("Calling gRPC createBlog Method");
        CreateBlogResponse blog = null;
        try {
            blog = syncClient.withDeadline(deadline)
                    .createBlog(constructRequest());

            Stream.of(blog).map(s -> s.getBlog()).forEach(this::responseFromServer);

        } catch (StatusRuntimeException e) {
            if (Status.INVALID_ARGUMENT.equals(e.getStatus())) {
                System.out.println("INVALID_ARGUMENT error");
            } else if (Status.DEADLINE_EXCEEDED.equals(e.getStatus())) {
                System.out.println("DEADLINE_EXCEEDED error");
            } else {
                e.printStackTrace();
            }
        }

        System.out.println("Calling gRPC readBlog Method");
        final ReadBlogResponse readBlogResponse = syncClient.withDeadline(deadline).readBlog(
                ReadBlogRequest.newBuilder()
                        .setBlogId(blog.getBlog().getId())
                        .build());
        Stream.of(readBlogResponse).map(s -> s.getBlog()).forEach(this::responseFromServer);

        final boolean testNotCreatedMongoDocument = true;
        if (testNotCreatedMongoDocument) {
            try {
                ReadBlogResponse readBlogResponseNotFound = syncClient.withDeadline(deadline).readBlog(
                        ReadBlogRequest.newBuilder()
                                .setBlogId("5d753f98549d1573a583ebbb")
                                .build());

                Stream.of(readBlogResponseNotFound).map(s -> s.getBlog()).forEach(this::responseFromServer);
            } catch (StatusRuntimeException e) {
                if (Status.INVALID_ARGUMENT.equals(e.getStatus())) {
                    System.out.println("INVALID_ARGUMENT error");
                } else if (Status.DEADLINE_EXCEEDED.equals(e.getStatus())) {
                    System.out.println("DEADLINE_EXCEEDED error");
                } else {
                    e.printStackTrace();
                }
            }
        }


        System.out.println("Calling gRPC updateBlog Method");

        UpdateBlogRequest updateBlogRequest = UpdateBlogRequest.newBuilder()
                .setBlog(Blog.newBuilder()
                        .setId(blog.getBlog().getId())
                        .setAuthorId("Updated Author")
                        .setContent("Updated Content created: " + Stopwatch.createStarted())
                        .setTitle("Updated Title 1")
                        .build())
                .build();
        try {
            Stream.of(syncClient.updateBlog(updateBlogRequest)).map(s -> s.getBlog()).forEach(this::responseFromServer);
        } catch (StatusRuntimeException e) {
            if (Status.INVALID_ARGUMENT.equals(e.getStatus())) {
                System.out.println("INVALID_ARGUMENT error");
            } else if (Status.DEADLINE_EXCEEDED.equals(e.getStatus())) {
                System.out.println("DEADLINE_EXCEEDED error");
            } else {
                e.printStackTrace();
            }
        }


        System.out.println("Calling gRPC deleteBlog Method");
        final DeleteBlogResponse deleteBlogResponse = syncClient.deleteBlog(DeleteBlogRequest.newBuilder().setBlogId(blog.getBlog().getId()).build());
        Stream.of(deleteBlogResponse).map(s -> s.getBlogId()).forEach(System.out::println);


        System.out.println("Calling gRPC listBlog Method");
        Streams.stream(syncClient.listBlog(ListBlogRequest.newBuilder().build()))
                .map(s -> s.getBlog())
                .forEach(this::responseFromServer);

        System.out.println("Shutting down channel");
        Uninterruptibles.sleepUninterruptibly(4000L, TimeUnit.MILLISECONDS);
        channel.shutdown();

    }

    private CreateBlogRequest constructRequest() {
        Blog blog = Blog.newBuilder()
                .setAuthorId("Author")
                .setContent("Content created: " + Stopwatch.createStarted())
                .setTitle("Title 1")
                .build();
        return CreateBlogRequest.newBuilder()
                .setBlog(blog)
                .build();
    }

    private void responseFromServer(Blog blog) {
        System.out.println("Response from server, ID:" + blog.getId() + " AuthorId:" + blog.getAuthorId() + " Content: " + blog.getContent());
    }

}
