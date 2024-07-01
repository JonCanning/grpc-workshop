import { GreeterClient, HelloRequest } from "./proto/hello";
import * as grpc from "@grpc/grpc-js";

const startClient = async (callback) => {
  const credentials = grpc.ChannelCredentials.createInsecure();
  const client = new GreeterClient("localhost:8080", credentials);
  const request = new HelloRequest();
  request.name = "TS";
  client.SayHello(request, (err, response) => {
    if (err) {
      console.error("Failed to call SayHello:", err);
      return;
    }
    console.log("Response:", response.message);
    callback();
  });
};

export { startClient };
