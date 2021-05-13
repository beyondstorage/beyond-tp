import 'package:flutter/material.dart';

class ConfirmTitle extends StatelessWidget {
  final String title;
  final IconData? icon;

  ConfirmTitle({
    required this.title,
    required this.icon,
  });

  @override
  Widget build(BuildContext context) {
    return Flex(
      direction: Axis.horizontal,
      verticalDirection: VerticalDirection.up,
      children: [
        Icon(this.icon, size: 24.0),
        Padding(
          padding: EdgeInsets.only(left: 12.0),
          child: SelectableText(
            this.title,
            maxLines: 1,
            style: Theme.of(context).textTheme.headline5
          ),
        ),
      ],
    );
  }
}