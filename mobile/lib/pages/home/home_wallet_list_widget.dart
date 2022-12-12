import 'package:flutter/material.dart';
import 'package:flutter/src/widgets/container.dart';
import 'package:flutter/src/widgets/framework.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:walletup/common/data.dart';
import 'package:walletup/pages/home/home_page.dart';
import 'package:walletup/pages/wallet_list/wallet_list_page.dart';
import 'package:walletup/widgets/wallet_card_widget.dart';

class HomeWalletListWidget extends StatefulWidget {
  const HomeWalletListWidget({super.key});

  @override
  State<HomeWalletListWidget> createState() => HomeWalletListWidgetState();
}

class HomeWalletListWidgetState extends State<HomeWalletListWidget> {
  @override
  Widget build(BuildContext context) {
    const double marginHorizontal = 16;

    return Container(
      margin: const EdgeInsets.symmetric(horizontal: marginHorizontal),
      child: Column(
        children: [
          Row(
            children: [
              Text(
                "Suas carteiras",
                style: GoogleFonts.poppins(color: Colors.white, fontSize: 15),
              ),
              const Spacer(),
              GestureDetector(
                onTap: () {
                  Navigator.push(
                      context,
                      MaterialPageRoute(
                          builder: (context) => const WalletListPage()));
                },
                child: Text(
                  "Ver todas",
                  style: GoogleFonts.poppins(
                    color: Colors.grey,
                    fontWeight: FontWeight.w300,
                    fontSize: 10,
                  ),
                ),
              ),
            ],
          ),
          Container(
            margin: const EdgeInsets.only(top: 20),
            height: 400,
            child: ListView(
              scrollDirection: Axis.horizontal,
              children: [
                for (var w in walletsTemplate) WalletCardWidget(wallet: w),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
