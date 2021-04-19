import 'package:flutter/material.dart';

import '../../common/colors.dart';

class TaskStatus extends StatelessWidget {
  final String status;

  Color get bgColor {
    switch (status) {
      case "created":
        return rgba(255, 206, 52, 1);
      case "stopped":
        return rgba(202, 38, 33, 1);
      case "finished":
        return rgba(0, 170, 114, 1);
      default:
        return rgba(0, 170, 114, 1);
    }
  }

  IconData get iconName {
    switch (status) {
      case "created":
        return Icons.more_horiz;
      case "stopped":
        return Icons.pause;
      case "finished":
        return Icons.done;
      default:
        return Icons.done;
    }
  }

  TaskStatus(this.status);

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.symmetric(vertical: 12.0, horizontal: 16.0),
      child: Wrap(
        crossAxisAlignment: WrapCrossAlignment.center,
        spacing: 6,
        children: [
          Container(
            width: 12,
            height: 12,
            alignment: Alignment.center,
            decoration: BoxDecoration(
              color: bgColor,
              borderRadius: BorderRadius.circular(6.0),
            ),
            child: Icon(
              iconName,
              size: 8,
              color: Colors.white,
            ),
          ),
          SelectableText(
            status,
            style: Theme.of(context).textTheme.bodyText1,
          ),
        ],
      ),
    );
  }
}
