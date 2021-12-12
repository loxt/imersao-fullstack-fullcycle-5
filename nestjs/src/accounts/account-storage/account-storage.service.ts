import { Injectable, Scope } from '@nestjs/common';
import { Account } from '../entities/account.entity';
import { AccountsService } from '../accounts.service';

@Injectable({
  scope: Scope.REQUEST,
})
export class AccountStorageService {
  private _account?: Account = null;

  constructor(private readonly accountsService: AccountsService) {}

  get account() {
    return this._account;
  }

  async setBy(token: string) {
    this._account = await this.accountsService.findOne(token);
  }
}
