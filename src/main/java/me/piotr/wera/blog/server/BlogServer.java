package me.piotr.wera.blog.server;

import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoClients;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoDatabase;
import com.mongodb.client.result.DeleteResult;
import io.grpc.Server;
import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import io.vavr.Predicates;
import me.piotr.wera.*;
import me.piotr.wera.common.GrpcCommons;
import me.piotr.wera.common.MongoCommons;
import org.bson.Document;
import org.bson.types.ObjectId;

import java.util.Objects;
import java.util.function.Consumer;
import java.util.stream.Collectors;
import java.util.stream.Stream;

import static com.mongodb.client.model.Filters.eq;

public class BlogServer extends BlogServiceGrpc.BlogServiceImplBase implements GrpcCommons, MongoCommons {

    private final MongoClient mongoClient = MongoClients.create("mongodb://root:example@localhost:27017");
    private final MongoDatabase database = mongoClient.getDatabase("mydb");
    private final MongoCollection<Document> collection = database.getCollection("blog");

    public static void main(String[] args) throws Exception {
        new BlogServer().go();
    }

    public void go() throws Exception {
        final Server server = createServer(USE_SSL);
        server.start();
        System.out.println("Server listening for connection on port " + GRPC_SERVER_PORT);
        Runtime.getRuntime().addShutdownHook(new Thread(() -> onServerShutdown(server)));
        server.awaitTermination();
    }

    @Override
    public void createBlog(CreateBlogRequest request, StreamObserver<CreateBlogResponse> responseObserver) {
        Stream.of("createBlog function", this).forEach(System.out::println);

        final Blog blog = request.getBlog();
        final Document document = mapBlogToDocument(blog);

        collection.insertOne(document);

        final String id = document.getObjectId("_id").toString();

        final CreateBlogResponse response = CreateBlogResponse.newBuilder()
                .setBlog(Blog.newBuilder()
                        .setId(id)
                        .setAuthorId(blog.getAuthorId())
                        .setContent(blog.getContent())
                        .setTitle(blog.getTitle()))
                .build();

        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void readBlog(ReadBlogRequest request, StreamObserver<ReadBlogResponse> responseObserver) {
        Stream.of("readBlog function", this).forEach(System.out::println);

        final String blogId = request.getBlogId();

        Consumer<Document> operateOnDocument = (document) -> {
            if (document == null) {
                responseObserver.onError(Status.NOT_FOUND.withDescription("The blog with the corresponding id was not found").asRuntimeException());
                return;
            }
            final ReadBlogResponse response = constuctReadBlogResponse(document, blogId);
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        };

        Stream.of(blogId)
                .map(this::parseObjectId)
                .filter(Predicates.not(Objects::isNull))
                .map(o -> collection.find(eq("_id", o)).first())
                .collect(Collectors.toList())
                .forEach(operateOnDocument);
    }

    @Override
    public void updateBlog(UpdateBlogRequest request, StreamObserver<UpdateBlogResponse> responseObserver) {
        Stream.of("updateBlog function", this).forEach(System.out::println);

        final String blogId = request.getBlog().getId();


        final ObjectId objectId = parseObjectId(blogId);
        if (objectId == null) {
            responseObserver.onError(Status.INVALID_ARGUMENT.withDescription("Couldnt parse blogid").asRuntimeException());
            return;
        }
        final Document document = collection.find(eq("_id", objectId)).first();
        if (document == null) {
            responseObserver.onError(Status.NOT_FOUND.withDescription("The blog with the corresponding id was not found").asRuntimeException());
            return;
        }
        final Document replacement = mapBlogToDocument(request.getBlog());
        collection.replaceOne(eq("_id", document.getObjectId("_id")), replacement);
        final UpdateBlogResponse response = constuctUpdateBlogResponse(replacement, blogId);
        responseObserver.onNext(response);
        responseObserver.onCompleted();

    }

    @Override
    public void deleteBlog(DeleteBlogRequest request, StreamObserver<DeleteBlogResponse> responseObserver) {
        Stream.of("deleteBlog function", this).forEach(System.out::println);

        final String blogId = request.getBlogId();


        final ObjectId objectId = parseObjectId(blogId);
        if (objectId == null) {
            responseObserver.onError(Status.INVALID_ARGUMENT.withDescription("Couldn't parse blogid").asRuntimeException());
            return;
        }

        DeleteResult deleteResult = null;
        try {
            deleteResult = collection.deleteOne(eq("_id", objectId));
        } catch (Exception e) {
            throwError(responseObserver, e, Status.INVALID_ARGUMENT, "Couldnt parse blogid");
        }

        if (deleteResult.getDeletedCount() == 0) {
            throwError(responseObserver, null, Status.INVALID_ARGUMENT, "Couldnt parse blogid");
            return;
        }
        responseObserver.onNext(DeleteBlogResponse.newBuilder().setBlogId(blogId).build());
        responseObserver.onCompleted();

    }

    @Override
    public void listBlog(ListBlogRequest request, StreamObserver<ListBlogResponse> responseObserver) {
        collection.find().iterator().forEachRemaining(document -> responseObserver.onNext(ListBlogResponse.newBuilder().setBlog(documentToBlog(document)).build()));
        responseObserver.onCompleted();
    }


    private ObjectId parseObjectId(String blogId) {
        try {
            return new ObjectId(blogId);
        } catch (IllegalArgumentException e) {
            return null;
        }
    }
    private Blog documentToBlog(Document document) {
        return Blog.newBuilder()
                .setAuthorId(document.getString("author_id"))
                .setContent(document.getString("content"))
                .setTitle(document.getString("title")).build();
    }

    private ReadBlogResponse constuctReadBlogResponse(Document document, String blogId) {
        return ReadBlogResponse.newBuilder()
                .setBlog(Blog.newBuilder()
                        .setId(blogId)
                        .setAuthorId(document.getString("author_id"))
                        .setContent(document.getString("content"))
                        .setTitle(document.getString("title")))
                .build();
    }

    private UpdateBlogResponse constuctUpdateBlogResponse(Document document, String blogId) {
        return UpdateBlogResponse.newBuilder()
                .setBlog(Blog.newBuilder()
                        .setId(blogId)
                        .setAuthorId(document.getString("author_id"))
                        .setContent(document.getString("content"))
                        .setTitle(document.getString("title")))
                .build();
    }

    private Document mapBlogToDocument(Blog blog) {
        final Document document = new Document("author_id", blog.getAuthorId());
        document.append("title", blog.getTitle());
        document.append("content", blog.getContent());
        return document;

    }

    private <T> void throwError(StreamObserver<T> responseObserver, Throwable throwable, Status status, String message) {
        Status status1 = status
                .withDescription(message);
        if (throwable != null) {
            status1.augmentDescription(throwable.getLocalizedMessage());
        }
        responseObserver.onError(status1.asRuntimeException());
    }

}
