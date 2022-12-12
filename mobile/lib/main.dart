import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:walletup/common/colors.dart';
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
    return MaterialApp(
      home: const HomePage(),
      theme: ThemeData(
        // fontFamily: 'poppins',
        primaryColor: green,
        backgroundColor: Colors.black,
        colorScheme: const ColorScheme.dark(),
        scaffoldBackgroundColor: Colors.black,
        brightness: Brightness.dark,

        // FONTS
        textTheme: TextTheme(
          titleLarge: GoogleFonts.poppins(),
        ),
      ),
    );
  }
}
