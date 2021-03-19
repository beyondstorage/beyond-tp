import 'package:flutter/material.dart';

import 'mobile.dart';
import 'desktop.dart';

class Layout extends StatelessWidget {
  final Widget child;

  Layout({ this.child });

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(builder: (context, constraints) {
      if (constraints.maxWidth < 600) {
        return MobileLayout(child: child);
      }

      return DesktopLayout(child: child);
    });
  }
}