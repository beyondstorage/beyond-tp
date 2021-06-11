import 'package:flutter/material.dart';

// import 'header.dart';
import 'navigators.dart';

class DesktopLayout extends StatelessWidget {
  final Widget child;

  DesktopLayout({ required this.child });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      // appBar: Header(),
      body: Row(
        children: [
          Navigators(),
          Expanded(
            child: Padding(
              child: child,
              padding: EdgeInsets.all(20.0),
            ),
          ),
        ],
      ),
    );
  }
}