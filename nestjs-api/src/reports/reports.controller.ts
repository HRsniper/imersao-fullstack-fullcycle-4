import { Body, Controller, Get, Post, UseGuards } from "@nestjs/common";
import { JwtAuthGuard } from "../auth/jwt-auth.guard";
import { TenantGuard } from "../tenants/tenant.guard";
import { CreateReportDto } from "./dto/create-report.dto";
import { ReportsService } from "./reports.service";

@UseGuards(JwtAuthGuard, TenantGuard)
@Controller("reports")
export class ReportsController {
  constructor(private readonly reportsService: ReportsService) {}

  @Post()
  create(@Body() createReportDto: CreateReportDto) {
    return this.reportsService.create(createReportDto);
  }

  @Get()
  findAll() {
    return this.reportsService.findAll();
  }
}
