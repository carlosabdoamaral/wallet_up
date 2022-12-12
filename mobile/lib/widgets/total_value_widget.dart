import 'package:flutter/material.dart';
import 'package:flutter/src/widgets/container.dart';
import 'package:flutter/src/widgets/framework.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:walletup/common/variables.dart';

class TotalValueWidget extends StatefulWidget {
  const TotalValueWidget(
      {super.key, required this.value, required this.currency});

  final double value;
  final String currency;

  @override
  State<TotalValueWidget> createState() => _TotalValueWidgetState();
}

class _TotalValueWidgetState extends State<TotalValueWidget> {
  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      alignment: Alignment.center,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.center,
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Text(
            widget.value.toString(),
            style: GoogleFonts.poppins(
              color: Colors.white,
              fontSize: 35,
              fontWeight: FontWeight.w500,
            ),
          ),
          Text(
            widget.currency,
            style: GoogleFonts.poppins(
              color: Colors.grey,
              fontSize: 15,
            ),
          ),
        ],
      ),
    );
  }
}
