import 'dart:ffi';

class SharedContactModel {
  SharedContactModel(this.id, this.email);
  final int id;
  final String email;
}

class CategoryModel {
  CategoryModel(this.id, this.title);
  final int id;
  final String title;
}

class TransactionModel {
  TransactionModel(
    this.id,
    this.accountID,
    this.value,
    this.currencyID,
    this.categoryID,
    this.wallet,
    this.description,
    this.sentAt
  );

  final int id;
  final int accountID;
  final double value;
  final int currencyID;
  final CategoryModel categoryID;
  final WalletBasicModel wallet;
  final String description;
  final DateTime sentAt;
}

class WalletBasicModel {
  WalletBasicModel(this.id, this.title, this.description);
  final int id;
  final String title;
  final String description;
}

class WalletModel {
  WalletModel(
    this.id,
    this.accountID,
    this.currencyID,
    this.title,
    this.description,
    this.sharedContacts,
    this.transactionList,
    this.createdAt,
    this.updatedAt
  );

  final int id;
  final int accountID;
  final int currencyID;
  final String title;
  final String description;
  final List<SharedContactModel> sharedContacts;
  final List<TransactionModel> transactionList;
  final DateTime createdAt;
  final DateTime updatedAt;
}
