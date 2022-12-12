import 'package:flutter/material.dart';
import 'package:flutter/src/widgets/container.dart';
import 'package:flutter/src/widgets/framework.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:walletup/common/colors.dart';
import 'package:walletup/pages/wallet_list/wallet_list_page.dart';
import 'package:walletup/widgets/total_value_widget.dart';

class WalletDetailsPage extends StatefulWidget {
  const WalletDetailsPage({super.key});

  @override
  State<WalletDetailsPage> createState() => _WalletDetailsPageState();
}

class _WalletDetailsPageState extends State<WalletDetailsPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.black,
        actions: [
          Container(
            margin: EdgeInsets.only(right: 16),
            child: Icon(
              Icons.add,
              size: 28,
            ),
          ),
        ],
      ),
      body: ListView(
        children: [
          SafeArea(
            child: Container(
              margin: EdgeInsets.symmetric(horizontal: 16),
              child: Column(
                children: [
                  Container(
                    width: double.infinity,
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text(
                          "Viagem xyz",
                          overflow: TextOverflow.ellipsis,
                          style: GoogleFonts.poppins(
                            fontSize: 24,
                            fontWeight: FontWeight.w500,
                          ),
                        ),
                        Text(
                          "descricao da viagem xyz",
                          textAlign: TextAlign.justify,
                          maxLines: 5,
                          style: GoogleFonts.poppins(
                            color: Colors.white60,
                            fontWeight: FontWeight.w300,
                          ),
                        ),
                      ],
                    ),
                  ),
                  const SizedBox(height: 50),
                  const TotalValueWidget(value: 4000, currency: "USD"),
                  const SizedBox(height: 50),
                  Row(
                    children: [
                      Text(
                        "Compartilhado com",
                        style: GoogleFonts.poppins(
                          color: Colors.white,
                          fontSize: 15,
                        ),
                      ),
                      const Spacer(),
                      GestureDetector(
                        onTap: () {
                          print("SHARE");
                        },
                        child: Text(
                          "Adicionar",
                          style: GoogleFonts.poppins(
                            color: Colors.grey,
                            fontWeight: FontWeight.w300,
                            fontSize: 10,
                          ),
                        ),
                      ),
                    ],
                  ),
                  SizedBox(
                    height: 90,
                    child: ListView(
                      scrollDirection: Axis.horizontal,
                      children: const [
                        CircleAvatar(
                            backgroundColor: Colors.white, radius: 35.0),
                        SizedBox(width: 8),
                        CircleAvatar(
                            backgroundColor: Colors.white, radius: 35.0),
                        SizedBox(width: 8),
                        CircleAvatar(
                            backgroundColor: Colors.white, radius: 35.0),
                        SizedBox(width: 8),
                        CircleAvatar(
                            backgroundColor: Colors.white, radius: 35.0),
                        SizedBox(width: 8),
                        CircleAvatar(
                            backgroundColor: Colors.white, radius: 35.0),
                        SizedBox(width: 8),
                        CircleAvatar(
                            backgroundColor: Colors.white, radius: 35.0),
                        SizedBox(width: 8),
                        CircleAvatar(
                            backgroundColor: Colors.white, radius: 35.0),
                        SizedBox(width: 8),
                        CircleAvatar(
                            backgroundColor: Colors.white, radius: 35.0),
                        SizedBox(width: 8),
                        CircleAvatar(
                            backgroundColor: Colors.white, radius: 35.0),
                        SizedBox(width: 8),
                        CircleAvatar(
                            backgroundColor: Colors.white, radius: 35.0),
                        SizedBox(width: 8),
                      ],
                    ),
                  ),
                  const SizedBox(
                    height: 60,
                  ),
                  Row(
                    children: [
                      Text(
                        "Transações",
                        style: GoogleFonts.poppins(
                          color: Colors.white,
                          fontSize: 15,
                        ),
                      ),
                      const Spacer(),
                      GestureDetector(
                        onTap: () {
                          print("SHARE");
                        },
                        child: Text(
                          "Adicionar",
                          style: GoogleFonts.poppins(
                            color: Colors.grey,
                            fontWeight: FontWeight.w300,
                            fontSize: 10,
                          ),
                        ),
                      ),
                    ],
                  ),
                  Column(
                    children: [
                      Container(
                        margin: const EdgeInsets.symmetric(vertical: 8.0),
                        child: Expanded(
                          child: Container(
                            padding: const EdgeInsets.all(16.0),
                            decoration: const BoxDecoration(
                              borderRadius:
                                  BorderRadius.all(Radius.circular(5)),
                              color: darkGrey,
                            ),
                            child: Row(
                              children: [
                                Column(
                                  crossAxisAlignment: CrossAxisAlignment.start,
                                  children: const [
                                    Text("22/11/2022"),
                                    Text("Compras no shopping"),
                                  ],
                                ),
                                const Spacer(),
                                const Text("-100"),
                              ],
                            ),
                          ),
                        ),
                      ),
                    ],
                  ),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}
