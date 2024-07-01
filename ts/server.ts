import { UnimplementedGreeterService, HelloReply } from "./proto/hello";
import * as grpc from "@grpc/grpc-js";

const greeterService: UnimplementedGreeterService = {
  SayHello: (call, callback) => {
    const name = call.request.name;
    console.log(`Received request from ${name}`);
    const reply = new HelloReply();
    reply.message = `Hello ${name} from TS`;
    callback(null, reply);
  },
};

const startServer = async () => {
  const server = new grpc.Server();
  server.addService(UnimplementedGreeterService.definition, greeterService);
  server.bindAsync(
    "localhost:8080",
    grpc.ServerCredentials.createInsecure(),
    (err, _) => {
      console.log("Server started");
      if (err) {
        console.error("Failed to bind server:", err);
        return;
      }
    }
  );
};

export { startServer };
