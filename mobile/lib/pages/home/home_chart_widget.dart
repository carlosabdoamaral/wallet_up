import 'package:flutter/material.dart';
import 'package:flutter/src/widgets/container.dart';
import 'package:flutter/src/widgets/framework.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:syncfusion_flutter_charts/charts.dart';
import 'package:walletup/common/colors.dart';
import 'package:walletup/common/data.dart';
import 'package:walletup/common/models.dart';
import 'package:walletup/common/variables.dart';
import 'package:walletup/utils/init_variables.dart';

class HomeChartWidget extends StatefulWidget {
  const HomeChartWidget({super.key});

  @override
  State<HomeChartWidget> createState() => _HomeChartWidgetState();
}

class _HomeChartWidgetState extends State<HomeChartWidget> {
  List<HomeChartModel> data = [];

  @override
  void initState() {
    super.initState();
    filterData();
  }

  filterData() {
    for (var t in getTransactions()) {
      setState(() {
        data.add(HomeChartModel(t.sentAt, t.value));
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: 80,
      child: SfCartesianChart(
        plotAreaBorderColor: Colors.transparent,
        primaryXAxis: DateTimeAxis(
          axisLine: const AxisLine(width: 0),
          majorGridLines: const MajorGridLines(width: 0),
          isVisible: false,
          labelStyle: const TextStyle(fontSize: 0),
        ),
        primaryYAxis: DateTimeAxis(
          axisLine: const AxisLine(width: 0),
          majorGridLines: const MajorGridLines(width: 0),
          isVisible: false,
          labelStyle: const TextStyle(fontSize: 0),
        ),
        borderColor: Colors.transparent,
        series: <ChartSeries>[
          LineSeries<HomeChartModel, DateTime>(
            color: balanceValue > 0 ? green : red,
            dataSource: data,
            xValueMapper: (HomeChartModel k, _) => k.time,
            yValueMapper: (HomeChartModel k, _) => k.value,
          )
        ],
      ),
    );
  }
}
