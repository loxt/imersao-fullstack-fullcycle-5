import { Module } from '@nestjs/common';
import { OrdersModule } from './orders/orders.module';
import { SequelizeModule } from '@nestjs/sequelize';
import { Order } from './orders/entities/order.entity';
import { AccountsModule } from './accounts/accounts.module';
import { Account } from './accounts/entities/account.entity';
import { ConfigModule } from '@nestjs/config';

@Module({
  imports: [
    OrdersModule,
    ConfigModule.forRoot(),
    SequelizeModule.forRoot({
      dialect: process.env.DB_CONNECTION as any,
      host: process.env.DB_HOST,
      port: parseInt(process.env.DB_PORT),
      username: process.env.DB_USERNAME,
      password: process.env.DB_PASSWORD,
      database: process.env.DB_DATABASE,
      autoLoadModels: true,
      models: [Order, Account],
      sync: {
        alter: true,
        //force: true
      },
    }),
    // SequelizeModule.forRoot({
    //   dialect: 'sqlite',
    //   host: join(__dirname, 'database.sqlite'),
    //   autoLoadModels: true,
    //   models: [Order, Account],
    //   sync: {
    //     // alter: true,
    //     // force: true,
    //   },
    // }),
    AccountsModule,
  ],
})
export class AppModule {}
