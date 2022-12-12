import 'package:flutter/material.dart';
import 'package:flutter/src/widgets/container.dart';
import 'package:flutter/src/widgets/framework.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:walletup/common/colors.dart';
import 'package:walletup/pages/wallet_details/wallet_details_page.dart';

class WalletListPage extends StatefulWidget {
  const WalletListPage({super.key});

  @override
  State<WalletListPage> createState() => _WalletListPageState();
}

class _WalletListPageState extends State<WalletListPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        centerTitle: true,
        backgroundColor: Colors.black,
        title: Text("Carteiras"),
        actions: [
          Container(
            margin: EdgeInsets.only(right: 16),
            child: Icon(
              Icons.add,
            ),
          ),
        ],
      ),
      body: ListView(
        children: const [
          WalletListRowWidget(),
          WalletListRowWidget(),
          WalletListRowWidget(),
          WalletListRowWidget(),
          WalletListRowWidget(),
          WalletListRowWidget(),
          WalletListRowWidget(),
          WalletListRowWidget(),
        ],
      ),
    );
  }
}

class WalletListRowWidget extends StatelessWidget {
  const WalletListRowWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {
        Navigator.push(context,
            MaterialPageRoute(builder: (context) => const WalletDetailsPage()));
      },
      child: Container(
        margin: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
        padding: const EdgeInsets.all(13.0),
        decoration: const BoxDecoration(
          borderRadius: BorderRadius.all(Radius.circular(8.0)),

          color: darkGrey,
          // color: const ColorScheme.dark().background,
        ),
        child: Row(
          children: [
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  "22/11/2022",
                  style: GoogleFonts.poppins(
                    color: Colors.white38,
                    fontSize: 12,
                  ),
                ),
                Text(
                  "Viagem de buenos",
                  style: GoogleFonts.poppins(
                    color: Colors.white70,
                    fontSize: 16,
                  ),
                ),
              ],
            ),
            const Spacer(),
            const Icon(
              Icons.chevron_right_rounded,
              color: Colors.white70,
            )
          ],
        ),
      ),
    );
  }
}
