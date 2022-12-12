import 'package:flutter/material.dart';
import 'package:flutter/src/widgets/container.dart';
import 'package:flutter/src/widgets/framework.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:walletup/common/colors.dart';
import 'package:walletup/common/data.dart';
import 'package:walletup/common/models.dart';
import 'package:walletup/widgets/transaction_row_widget.dart';

class WalletCardWidget extends StatefulWidget {
  const WalletCardWidget({super.key, required this.wallet});

  final WalletModel wallet;

  @override
  State<WalletCardWidget> createState() => _WalletCardWidgetState();
}

class _WalletCardWidgetState extends State<WalletCardWidget> {
  List<TransactionModel> transactions = [];

  @override
  void initState() {
    super.initState();
    var i = 0;
    for (var t in widget.wallet.transactionList) {
      if (i < 4) {
        setState(() {
          transactions.add(t);
        });
      } else {
        break;
      }

      i++;
    }
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 300,
      decoration: const BoxDecoration(
        color: darkGrey,
      ),
      margin: const EdgeInsets.all(2.0),
      child: Container(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          children: [
            Container(
              margin: const EdgeInsets.only(bottom: 20),
              child: Column(
                children: [
                  Text(
                    widget.wallet.title,
                    textAlign: TextAlign.center,
                    style: GoogleFonts.poppins(
                      color: Colors.white,
                      fontSize: 18,
                    ),
                  ),
                  Text(
                    widget.wallet.title,
                    textAlign: TextAlign.center,
                    style: GoogleFonts.poppins(
                      color: Colors.grey,
                      fontSize: 10,
                    ),
                  ),
                ],
              ),
            ),
            // for (var i = 0; i <= 4; i++)
            //   for (var t in widget.wallet.transactionList)
            for (var t in transactions) TransactionRowWidget(transaction: t),
            const Spacer(),
            Center(
              child: GestureDetector(
                child: Container(
                  margin: EdgeInsets.only(bottom: 20),
                  child: Text(
                    "Adicionar transação",
                    style: GoogleFonts.poppins(
                      color: Colors.white,
                      fontSize: 12,
                    ),
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
