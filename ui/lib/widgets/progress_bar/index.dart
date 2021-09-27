import 'package:flutter/material.dart';

import '../../common/colors.dart';

class ProgressBar extends StatelessWidget {
  final double ratio;
  final double barWidth;
  final double barHeight;
  final String description;
  final Color ?barColor;

  ProgressBar({
    required this.ratio,
    this.barWidth = 140,
    this.barHeight = 6,
    this.description = '',
    this.barColor,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisSize: MainAxisSize.min,
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Row(
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            Container(
              width: barWidth,
              height: barHeight,
              alignment: Alignment.topLeft,
              decoration: BoxDecoration(
                color: barColor != null ? barColor?.withOpacity(0.1) : rgba(242, 242, 242, 1),
                borderRadius: BorderRadius.all(Radius.circular(30)),
              ),
              child: Container(
                width: ratio * barWidth,
                height: barHeight,
                decoration: BoxDecoration(
                  color: barColor ?? Theme.of(context).primaryColor,
                  borderRadius: BorderRadius.all(Radius.circular(30)),
                ),
              ),
            ),
            SizedBox(width: 8),
            Text(
              '${ratio * 100}%',
              style: TextStyle(fontSize: 12),
            )
          ],
        ),
        description.length > 0
            ? SelectableText(
                description,
                style: TextStyle(
                  fontSize: 12,
                  color: rgba(100, 116, 139, 1),
                ),
              )
            : Container(),
      ],
    );
  }
}
