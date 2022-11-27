import 'package:flutter/material.dart';
import 'package:flutter/src/widgets/container.dart';
import 'package:flutter/src/widgets/framework.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:walletup/common/variables.dart';

class HomeTotalValueWidget extends StatefulWidget {
  const HomeTotalValueWidget({super.key});

  @override
  State<HomeTotalValueWidget> createState() => _HomeTotalValueWidgetState();
}

class _HomeTotalValueWidgetState extends State<HomeTotalValueWidget> {
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
            balanceValue.toString(),
            style: GoogleFonts.poppins(
              color: Colors.white,
              fontSize: 35,
              fontWeight: FontWeight.w500,
            ),
          ),
          Text(
            "USD",
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
