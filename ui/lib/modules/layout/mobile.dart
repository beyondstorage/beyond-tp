import 'package:flutter/material.dart';

import 'header.dart';
import 'navigators.dart';

class MobileLayout extends StatelessWidget {
  final Widget child;

  MobileLayout({ required this.child });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: MobileHeader(),
      body: SingleChildScrollView(
        child: child,
        scrollDirection: Axis.vertical,
      ),
      drawer: Drawer(
        child: Navigators(),
      ),
    );
  }
}