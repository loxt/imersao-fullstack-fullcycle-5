import { Column, DataType, Model, Table } from 'sequelize-typescript';

@Table({
  tableName: 'accounts',
  createdAt: 'created_at',
  updatedAt: 'updated_at',
})
export class Account extends Model {
  @Column({
    type: DataType.UUID,
    defaultValue: DataType.UUIDV4,
    primaryKey: true,
  })
  id: string;

  @Column({
    allowNull: false,
    defaultValue: () => Math.random().toString(36).slice(2),
  })
  token: string;
}
