// WALLETs
import 'package:walletup/common/models.dart';

final List<WalletModel> walletsTemplate = [
  WalletModel(
    1,
    1,
    1,
    "Viagem de Buenos Aires",
    "Lorem ipsum dolor la la la",
    [
      SharedContactModel(1, "cab@gmail.com"),
    ],
    defaultTransactions,
    DateTime.parse("2022-07-01"),
    DateTime.parse("2022-07-01"),
  ),
];

final List<TransactionModel> defaultTransactions = [
  TransactionModel(
    1,
    1,
    100,
    1,
    CategoryModel(1, 'PROFIT'),
    WalletBasicModel(1, 'Teste', 'Abcdef'),
    'Viagem',
    DateTime.utc(2022, 11, 21),
  ),
  TransactionModel(
    1,
    1,
    100,
    1,
    CategoryModel(1, 'PROFIT'),
    WalletBasicModel(1, 'Teste', 'Abcdef'),
    'Viagem',
    DateTime.utc(2022, 11, 21),
  ),
  TransactionModel(
    1,
    1,
    250,
    1,
    CategoryModel(1, 'PROFIT'),
    WalletBasicModel(1, 'Teste', 'Abcdef'),
    'Viagem',
    DateTime.utc(2022, 11, 22),
  ),
  TransactionModel(
    1,
    1,
    -450,
    1,
    CategoryModel(1, 'PROFIT'),
    WalletBasicModel(1, 'Teste', 'Abcdef'),
    'Viagem',
    DateTime.utc(2022, 11, 23),
  ),
  TransactionModel(
    1,
    1,
    1000,
    1,
    CategoryModel(1, 'PROFIT'),
    WalletBasicModel(1, 'Teste', 'Abcdef'),
    'Viagem',
    DateTime.utc(2022, 11, 24),
  ),
];
