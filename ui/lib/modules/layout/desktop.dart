import 'package:flutter/material.dart';

import 'header.dart';
import 'navigators.dart';

class DesktopLayout extends StatelessWidget {
  final Widget child;

  DesktopLayout({ this.child });

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Column(
        mainAxisSize: MainAxisSize.max,
        children: [
          Header(),
          Expanded(
            child: Row(
              children: [
                Navigators(),
                Expanded(
                  child: Padding(
                    child: child,
                    padding: EdgeInsets.all(20.0),
                  ),
                ),
              ],
            )
          ),
        ],
      ),
      color: Theme.of(context).scaffoldBackgroundColor,
    );
  }
}