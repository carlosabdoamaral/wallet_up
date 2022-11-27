import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:flutter/src/widgets/container.dart';
import 'package:flutter/src/widgets/framework.dart';
import 'package:syncfusion_flutter_charts/charts.dart';
import 'package:walletup/common/colors.dart';
import 'package:walletup/common/data.dart';
import 'package:walletup/common/models.dart';
import 'package:walletup/pages/home/home_chart_widget.dart';
import 'package:walletup/pages/home/home_total_value_widget.dart';
import 'package:walletup/pages/home/home_wallet_list_widget.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.black,
      appBar: AppBar(
        title: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              "Carlos Amaral",
              style: GoogleFonts.poppins(
                color: Colors.white,
                fontSize: 20,
              ),
            ),
            Text(
              "carlosabdoamaral@gmail.com",
              style: GoogleFonts.poppins(
                color: Colors.grey,
                fontSize: 10,
              ),
            ),
          ],
        ),
        // backgroundColor: grey,
        backgroundColor: Colors.transparent,
        elevation: 0,
        toolbarHeight: 90,
        centerTitle: false,
      ),
      body: SafeArea(
        child: ListView(
          children: [
            Column(
              children: const [
                SizedBox(height: 40),
                HomeTotalValueWidget(),
                SizedBox(height: 20),
                HomeChartWidget(),
                SizedBox(height: 60),
                HomeWalletListWidget()
              ],
            ),
          ],
        ),
      ),
    );
  }
}
