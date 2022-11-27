import 'package:flutter/material.dart';
import 'package:flutter/src/widgets/container.dart';
import 'package:flutter/src/widgets/framework.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:walletup/common/colors.dart';
import 'package:walletup/common/data.dart';
import 'package:walletup/common/models.dart';
import 'package:walletup/pages/home/home_page.dart';
import 'package:walletup/widgets/wallet_card_widget.dart';

class TransactionRowWidget extends StatefulWidget {
  const TransactionRowWidget({super.key, required this.transaction});

  final TransactionModel transaction;

  @override
  State<TransactionRowWidget> createState() => _TransactionRowWidgetState();
}

class _TransactionRowWidgetState extends State<TransactionRowWidget> {
  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.symmetric(vertical: 4.0),
      padding: const EdgeInsets.all(10),
      decoration: const BoxDecoration(
        color: Color.fromRGBO(26, 25, 25, 1),
      ),
      child: Row(
        children: [
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                widget.transaction.sentAt.toUtc().toString(),
                style: GoogleFonts.poppins(
                  color: Colors.grey,
                  fontSize: 10,
                ),
              ),
              Text(
                widget.transaction.description,
                style: GoogleFonts.poppins(color: Colors.white70, fontSize: 12),
              ),
            ],
          ),
          const Spacer(),
          Text(
            widget.transaction.value.toString(),
            style: GoogleFonts.poppins(
                color: widget.transaction.value > 0 ? green : red,
                fontWeight: FontWeight.w600),
          ),
        ],
      ),
    );
  }
}
