import { Module } from "@nestjs/common";
import { SequelizeModule } from "@nestjs/sequelize";
import { AccountsController } from "./accounts.controller";
import { AccountsService } from "./accounts.service";
import { Account } from "./entities/account.entity";
import { MyAccountController } from "./my-account/my-account.controller";

@Module({
  imports: [SequelizeModule.forFeature([Account])],
  controllers: [AccountsController, MyAccountController],
  providers: [AccountsService]
})
export class AccountsModule {}
