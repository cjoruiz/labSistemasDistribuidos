package co.edu.unicauca.admin.proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.60.0)",
    comments = "Source: audio_streaming.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class StreamingServiceGrpc {

  private StreamingServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "servicios.StreamingService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<co.edu.unicauca.admin.proto.StreamRequest,
      co.edu.unicauca.admin.proto.AudioChunk> getStreamAudioMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "StreamAudio",
      requestType = co.edu.unicauca.admin.proto.StreamRequest.class,
      responseType = co.edu.unicauca.admin.proto.AudioChunk.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<co.edu.unicauca.admin.proto.StreamRequest,
      co.edu.unicauca.admin.proto.AudioChunk> getStreamAudioMethod() {
    io.grpc.MethodDescriptor<co.edu.unicauca.admin.proto.StreamRequest, co.edu.unicauca.admin.proto.AudioChunk> getStreamAudioMethod;
    if ((getStreamAudioMethod = StreamingServiceGrpc.getStreamAudioMethod) == null) {
      synchronized (StreamingServiceGrpc.class) {
        if ((getStreamAudioMethod = StreamingServiceGrpc.getStreamAudioMethod) == null) {
          StreamingServiceGrpc.getStreamAudioMethod = getStreamAudioMethod =
              io.grpc.MethodDescriptor.<co.edu.unicauca.admin.proto.StreamRequest, co.edu.unicauca.admin.proto.AudioChunk>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "StreamAudio"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  co.edu.unicauca.admin.proto.StreamRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  co.edu.unicauca.admin.proto.AudioChunk.getDefaultInstance()))
              .setSchemaDescriptor(new StreamingServiceMethodDescriptorSupplier("StreamAudio"))
              .build();
        }
      }
    }
    return getStreamAudioMethod;
  }

  private static volatile io.grpc.MethodDescriptor<co.edu.unicauca.admin.proto.AudioFileRequest,
      co.edu.unicauca.admin.proto.AudioFileResponse> getAlmacenarAudioMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AlmacenarAudio",
      requestType = co.edu.unicauca.admin.proto.AudioFileRequest.class,
      responseType = co.edu.unicauca.admin.proto.AudioFileResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<co.edu.unicauca.admin.proto.AudioFileRequest,
      co.edu.unicauca.admin.proto.AudioFileResponse> getAlmacenarAudioMethod() {
    io.grpc.MethodDescriptor<co.edu.unicauca.admin.proto.AudioFileRequest, co.edu.unicauca.admin.proto.AudioFileResponse> getAlmacenarAudioMethod;
    if ((getAlmacenarAudioMethod = StreamingServiceGrpc.getAlmacenarAudioMethod) == null) {
      synchronized (StreamingServiceGrpc.class) {
        if ((getAlmacenarAudioMethod = StreamingServiceGrpc.getAlmacenarAudioMethod) == null) {
          StreamingServiceGrpc.getAlmacenarAudioMethod = getAlmacenarAudioMethod =
              io.grpc.MethodDescriptor.<co.edu.unicauca.admin.proto.AudioFileRequest, co.edu.unicauca.admin.proto.AudioFileResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AlmacenarAudio"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  co.edu.unicauca.admin.proto.AudioFileRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  co.edu.unicauca.admin.proto.AudioFileResponse.getDefaultInstance()))
              .setSchemaDescriptor(new StreamingServiceMethodDescriptorSupplier("AlmacenarAudio"))
              .build();
        }
      }
    }
    return getAlmacenarAudioMethod;
  }

  private static volatile io.grpc.MethodDescriptor<co.edu.unicauca.admin.proto.CallbackRegistroRequest,
      co.edu.unicauca.admin.proto.CallbackRegistroResponse> getRegistrarCallbackMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RegistrarCallback",
      requestType = co.edu.unicauca.admin.proto.CallbackRegistroRequest.class,
      responseType = co.edu.unicauca.admin.proto.CallbackRegistroResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<co.edu.unicauca.admin.proto.CallbackRegistroRequest,
      co.edu.unicauca.admin.proto.CallbackRegistroResponse> getRegistrarCallbackMethod() {
    io.grpc.MethodDescriptor<co.edu.unicauca.admin.proto.CallbackRegistroRequest, co.edu.unicauca.admin.proto.CallbackRegistroResponse> getRegistrarCallbackMethod;
    if ((getRegistrarCallbackMethod = StreamingServiceGrpc.getRegistrarCallbackMethod) == null) {
      synchronized (StreamingServiceGrpc.class) {
        if ((getRegistrarCallbackMethod = StreamingServiceGrpc.getRegistrarCallbackMethod) == null) {
          StreamingServiceGrpc.getRegistrarCallbackMethod = getRegistrarCallbackMethod =
              io.grpc.MethodDescriptor.<co.edu.unicauca.admin.proto.CallbackRegistroRequest, co.edu.unicauca.admin.proto.CallbackRegistroResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RegistrarCallback"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  co.edu.unicauca.admin.proto.CallbackRegistroRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  co.edu.unicauca.admin.proto.CallbackRegistroResponse.getDefaultInstance()))
              .setSchemaDescriptor(new StreamingServiceMethodDescriptorSupplier("RegistrarCallback"))
              .build();
        }
      }
    }
    return getRegistrarCallbackMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static StreamingServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<StreamingServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<StreamingServiceStub>() {
        @java.lang.Override
        public StreamingServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new StreamingServiceStub(channel, callOptions);
        }
      };
    return StreamingServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static StreamingServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<StreamingServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<StreamingServiceBlockingStub>() {
        @java.lang.Override
        public StreamingServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new StreamingServiceBlockingStub(channel, callOptions);
        }
      };
    return StreamingServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static StreamingServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<StreamingServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<StreamingServiceFutureStub>() {
        @java.lang.Override
        public StreamingServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new StreamingServiceFutureStub(channel, callOptions);
        }
      };
    return StreamingServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void streamAudio(co.edu.unicauca.admin.proto.StreamRequest request,
        io.grpc.stub.StreamObserver<co.edu.unicauca.admin.proto.AudioChunk> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getStreamAudioMethod(), responseObserver);
    }

    /**
     */
    default void almacenarAudio(co.edu.unicauca.admin.proto.AudioFileRequest request,
        io.grpc.stub.StreamObserver<co.edu.unicauca.admin.proto.AudioFileResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAlmacenarAudioMethod(), responseObserver);
    }

    /**
     */
    default void registrarCallback(co.edu.unicauca.admin.proto.CallbackRegistroRequest request,
        io.grpc.stub.StreamObserver<co.edu.unicauca.admin.proto.CallbackRegistroResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRegistrarCallbackMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service StreamingService.
   */
  public static abstract class StreamingServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return StreamingServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service StreamingService.
   */
  public static final class StreamingServiceStub
      extends io.grpc.stub.AbstractAsyncStub<StreamingServiceStub> {
    private StreamingServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected StreamingServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new StreamingServiceStub(channel, callOptions);
    }

    /**
     */
    public void streamAudio(co.edu.unicauca.admin.proto.StreamRequest request,
        io.grpc.stub.StreamObserver<co.edu.unicauca.admin.proto.AudioChunk> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getStreamAudioMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void almacenarAudio(co.edu.unicauca.admin.proto.AudioFileRequest request,
        io.grpc.stub.StreamObserver<co.edu.unicauca.admin.proto.AudioFileResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAlmacenarAudioMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void registrarCallback(co.edu.unicauca.admin.proto.CallbackRegistroRequest request,
        io.grpc.stub.StreamObserver<co.edu.unicauca.admin.proto.CallbackRegistroResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRegistrarCallbackMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service StreamingService.
   */
  public static final class StreamingServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<StreamingServiceBlockingStub> {
    private StreamingServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected StreamingServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new StreamingServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public java.util.Iterator<co.edu.unicauca.admin.proto.AudioChunk> streamAudio(
        co.edu.unicauca.admin.proto.StreamRequest request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getStreamAudioMethod(), getCallOptions(), request);
    }

    /**
     */
    public co.edu.unicauca.admin.proto.AudioFileResponse almacenarAudio(co.edu.unicauca.admin.proto.AudioFileRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAlmacenarAudioMethod(), getCallOptions(), request);
    }

    /**
     */
    public co.edu.unicauca.admin.proto.CallbackRegistroResponse registrarCallback(co.edu.unicauca.admin.proto.CallbackRegistroRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRegistrarCallbackMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service StreamingService.
   */
  public static final class StreamingServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<StreamingServiceFutureStub> {
    private StreamingServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected StreamingServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new StreamingServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<co.edu.unicauca.admin.proto.AudioFileResponse> almacenarAudio(
        co.edu.unicauca.admin.proto.AudioFileRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAlmacenarAudioMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<co.edu.unicauca.admin.proto.CallbackRegistroResponse> registrarCallback(
        co.edu.unicauca.admin.proto.CallbackRegistroRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRegistrarCallbackMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_STREAM_AUDIO = 0;
  private static final int METHODID_ALMACENAR_AUDIO = 1;
  private static final int METHODID_REGISTRAR_CALLBACK = 2;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_STREAM_AUDIO:
          serviceImpl.streamAudio((co.edu.unicauca.admin.proto.StreamRequest) request,
              (io.grpc.stub.StreamObserver<co.edu.unicauca.admin.proto.AudioChunk>) responseObserver);
          break;
        case METHODID_ALMACENAR_AUDIO:
          serviceImpl.almacenarAudio((co.edu.unicauca.admin.proto.AudioFileRequest) request,
              (io.grpc.stub.StreamObserver<co.edu.unicauca.admin.proto.AudioFileResponse>) responseObserver);
          break;
        case METHODID_REGISTRAR_CALLBACK:
          serviceImpl.registrarCallback((co.edu.unicauca.admin.proto.CallbackRegistroRequest) request,
              (io.grpc.stub.StreamObserver<co.edu.unicauca.admin.proto.CallbackRegistroResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getStreamAudioMethod(),
          io.grpc.stub.ServerCalls.asyncServerStreamingCall(
            new MethodHandlers<
              co.edu.unicauca.admin.proto.StreamRequest,
              co.edu.unicauca.admin.proto.AudioChunk>(
                service, METHODID_STREAM_AUDIO)))
        .addMethod(
          getAlmacenarAudioMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              co.edu.unicauca.admin.proto.AudioFileRequest,
              co.edu.unicauca.admin.proto.AudioFileResponse>(
                service, METHODID_ALMACENAR_AUDIO)))
        .addMethod(
          getRegistrarCallbackMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              co.edu.unicauca.admin.proto.CallbackRegistroRequest,
              co.edu.unicauca.admin.proto.CallbackRegistroResponse>(
                service, METHODID_REGISTRAR_CALLBACK)))
        .build();
  }

  private static abstract class StreamingServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    StreamingServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return co.edu.unicauca.admin.proto.AudioStreaming.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("StreamingService");
    }
  }

  private static final class StreamingServiceFileDescriptorSupplier
      extends StreamingServiceBaseDescriptorSupplier {
    StreamingServiceFileDescriptorSupplier() {}
  }

  private static final class StreamingServiceMethodDescriptorSupplier
      extends StreamingServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    StreamingServiceMethodDescriptorSupplier(java.lang.String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (StreamingServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new StreamingServiceFileDescriptorSupplier())
              .addMethod(getStreamAudioMethod())
              .addMethod(getAlmacenarAudioMethod())
              .addMethod(getRegistrarCallbackMethod())
              .build();
        }
      }
    }
    return result;
  }
}
