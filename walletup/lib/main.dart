import 'package:flutter/material.dart';
import 'package:walletup/pages/home/home_page.dart';
import 'package:walletup/utils/init_variables.dart';

void main() {
  runApp(const Controller());
}

class Controller extends StatefulWidget {
  const Controller({super.key});

  @override
  State<Controller> createState() => _ControllerState();
}

class _ControllerState extends State<Controller> {
  @override
  void initState() {
    super.initState();
    initVariables();
  }

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      home: HomePage(),
    );
  }
}
