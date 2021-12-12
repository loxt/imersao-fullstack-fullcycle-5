import { Column, DataType, Model, Table } from 'sequelize-typescript';

export enum OrderStatus {
  Pending = 'pending',
  Approved = 'approved',
}

@Table({
  tableName: 'orders',
  createdAt: 'created_at',
  updatedAt: 'updated_at',
})
export class Order extends Model {
  @Column({
    type: DataType.UUIDV4,
    defaultValue: DataType.UUIDV4,
    primaryKey: true,
  })
  id: string;

  @Column({
    allowNull: false,
    type: DataType.DECIMAL(10, 2),
  })
  amount: number;

  @Column({
    allowNull: false,
  })
  credit_card_number: string;

  @Column({
    allowNull: false,
  })
  credit_card_name: string;

  @Column({
    allowNull: false,
  })
  status: OrderStatus;
}
