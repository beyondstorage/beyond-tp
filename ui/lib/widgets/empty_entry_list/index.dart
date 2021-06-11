import 'package:flutter/material.dart';
import 'package:ui/common/colors.dart';

import '../button/index.dart';
import '../button/constants.dart';
import '../page_container/index.dart';

class EmptyEntryList extends StatelessWidget {
  final IconData icon;
  final String title;
  final String subTitle;
  final String buttonText;
  final VoidCallback onClick;

  EmptyEntryList({
    required this.icon,
    required this.title,
    required this.subTitle,
    required this.buttonText,
    required this.onClick,
  });

  @override
  Widget build(BuildContext context) {
    return WidgetContainer(
      child: Container(
        height: 320,
        padding: EdgeInsets.only(top: 70),
        alignment: Alignment.center,
        child: Column(
          children: [
            Icon(icon, size: 56, color: defaultColor),
            SizedBox(height: 8),
            SelectableText(title, style: Theme.of(context).textTheme.headline6),
            SelectableText(
              subTitle,
              style: Theme.of(context).textTheme.bodyText2,
            ),
            SizedBox(height: 24),
            Button(
              icon: Icons.add,
              child: Text(buttonText),
              type: ButtonType.primary,
              onPressed: onClick,
            ),
          ],
        ),
      ),
    );
  }
}
