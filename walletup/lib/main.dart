import 'package:flutter/material.dart';
import 'package:walletup/home_page.dart';

void main() {
  runApp(const Controller());
}

class Controller extends StatelessWidget {
  const Controller({super.key});

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      home: HomePage(),
    );
  }
}