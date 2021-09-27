import { Body, Controller, Get, Post, Req, UseGuards } from "@nestjs/common";
import { JwtAuthGuard } from "../auth/jwt-auth.guard";
import { TenantGuard } from "../tenants/tenant.guard";
import { CreateTransactionDto } from "./dto/create-transaction.dto";
import { TransactionsService } from "./transactions.service";

@UseGuards(JwtAuthGuard, TenantGuard)
@Controller("transactions")
export class TransactionsController {
  constructor(private readonly transactionsService: TransactionsService) {}

  @Post()
  create(@Body() createTransactionDto: CreateTransactionDto) {
    return this.transactionsService.create(createTransactionDto);
  }

  @Get()
  findAll(@Req() req) {
    return this.transactionsService.findAll();
  }
}
