import 'package:walletup/common/data.dart';
import 'package:walletup/common/models.dart';
import 'package:walletup/common/variables.dart';

void initVariables() {
  setBalanceValue(getTransactions());
}

List<TransactionModel> getTransactions() {
  List<TransactionModel> res = [];

  // Faz o request
  for (var w in walletsTemplate) {
    for (var t in w.transactionList) {
      res.add(t);
    }
  }

  return res;
}

void setBalanceValue(List<TransactionModel> ts) {
  for (var t in ts) {
    balanceValue += t.value;
  }
}
