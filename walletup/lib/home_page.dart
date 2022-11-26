import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:flutter/src/widgets/container.dart';
import 'package:flutter/src/widgets/framework.dart';
import 'package:syncfusion_flutter_charts/charts.dart';
import 'package:walletup/common/colors.dart';
import 'package:walletup/common/data.dart';
import 'package:walletup/common/models.dart';

class SalesData {
  SalesData(this.year, this.sales);
  final DateTime year;
  final double sales;
}

class HomePage extends StatelessWidget {
  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    final List<SalesData> chartData = [
      SalesData(DateTime(2010), 35),
      SalesData(DateTime(2011), 28),
      SalesData(DateTime(2012), 34),
      SalesData(DateTime(2013), 32),
      SalesData(DateTime(2014), 40),
      SalesData(DateTime(2015), 35),
      SalesData(DateTime(2016), 28),
      SalesData(DateTime(2017), 34),
      SalesData(DateTime(2018), 32),
      SalesData(DateTime(2019), 40),
    ];

    final double _marginHorizontal = 16;

    return Scaffold(
      backgroundColor: Colors.black,
      appBar: AppBar(
        centerTitle: false,
        title: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              "Carlos Amaral",
              style: GoogleFonts.poppins(color: Colors.white, fontSize: 20),
            ),
            Text(
              "carlosabdoamaral@gmail.com",
              style: GoogleFonts.poppins(color: Colors.grey, fontSize: 10),
            ),
          ],
        ),
        backgroundColor: Colors.transparent,
        elevation: 0,
      ),
      body: SafeArea(
        child: ListView(
          children: [
            Column(
              children: [
                Container(
                  alignment: Alignment.center,
                  height: 150,
                  padding: EdgeInsets.only(top: 40),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.center,
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Text(
                        "2.400,50",
                        style: GoogleFonts.poppins(
                          color: Colors.white,
                          fontSize: 35,
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
                ),
                Container(
                  height: 150,
                  margin: EdgeInsets.only(bottom: 40),
                  child: SfCartesianChart(
                    primaryXAxis: DateTimeAxis(),
                    borderColor: Colors.transparent,
                    series: <ChartSeries>[
                      LineSeries<SalesData, DateTime>(
                          color: green,
                          dataSource: chartData,
                          xValueMapper: (SalesData sales, _) => sales.year,
                          yValueMapper: (SalesData sales, _) => sales.sales)
                    ],
                  ),
                ),
                Container(
                  margin: EdgeInsets.symmetric(horizontal: _marginHorizontal),
                  child: Column(
                    children: [
                      Row(
                        children: [
                          Text(
                            "Suas carteiras",
                            style: GoogleFonts.poppins(
                                color: Colors.white, fontSize: 15),
                          ),
                          const Spacer(),
                          Text(
                            "Ver todas",
                            style: GoogleFonts.poppins(
                              color: Colors.grey,
                              fontWeight: FontWeight.w300,
                              fontSize: 10,
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
                            for (var w in walletsTemplate)
                              WalletWidget(wallet: w),
                          ],
                        ),
                      ),
                    ],
                  ),
                )
              ],
            ),
          ],
        ),
      ),
    );
  }
}

class WalletWidget extends StatefulWidget {
  const WalletWidget({super.key, required this.wallet});

  final WalletModel wallet;

  @override
  State<WalletWidget> createState() => _WalletWidgetState();
}

class _WalletWidgetState extends State<WalletWidget> {
  @override
  Widget build(BuildContext context) {
    return Container(
      color: const Color.fromRGBO(30, 30, 30, 1),
      width: 300,
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
                    style: GoogleFonts.poppins(
                      color: Colors.white,
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
            for (var t in widget.wallet.transactionList)
              TransactionRow(
                transaction: t,
              ),
            const Spacer(),
            Center(
              child: Container(
                margin: EdgeInsets.only(bottom: 10),
                child: Text(
                  "Adicionar transação",
                  style: GoogleFonts.poppins(color: Colors.white, fontSize: 12),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}

class TransactionRow extends StatelessWidget {
  const TransactionRow({super.key, required this.transaction});
  final TransactionModel transaction;

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
                transaction.sentAt.toUtc().toString(),
                style: GoogleFonts.poppins(
                  color: Colors.grey,
                  fontSize: 10,
                ),
              ),
              Text(
                transaction.description,
                style: GoogleFonts.poppins(color: Colors.white70, fontSize: 12),
              ),
            ],
          ),
          const Spacer(),
          Text(
            transaction.value.toString(),
            style: GoogleFonts.poppins(
                color: transaction.value > 0 ? green : red,
                fontWeight: FontWeight.w600),
          ),
        ],
      ),
    );
  }
}
