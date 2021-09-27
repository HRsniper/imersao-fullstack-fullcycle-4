import { Injectable } from "@nestjs/common";

@Injectable()
export class AppService {
  isRunning(): string {
    return "Server is running!";
  }
}
