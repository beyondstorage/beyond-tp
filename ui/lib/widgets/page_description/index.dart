import 'package:flutter/material.dart';
import 'package:ui/common/colors.dart';

import '../page_container/index.dart';

class PageDescription extends StatelessWidget {
  final String title;
  final String? subtitle;
  final IconData? icon;
  final List<Widget>? children;

  PageDescription({
    this.icon = Icons.home,
    this.subtitle,
    this.children,
    required this.title,
  });

  @override
  Widget build(BuildContext context) {
    return WidgetContainer(
      margin: EdgeInsets.only(bottom: 16.0),
      child: Padding(
        padding: EdgeInsets.all(20.0),
        child: Row(
          children: [
            Container(
              width: 60.0,
              height: 60.0,
              margin: EdgeInsets.only(right: 16.0),
              decoration: new BoxDecoration(
                color: rgba(104, 131, 237, 0.1),
                borderRadius: BorderRadius.only(
                  topLeft: Radius.circular(60.0),
                  bottomLeft: Radius.circular(60.0),
                  bottomRight: Radius.circular(60.0),
                ),
              ),
              child: Center(
                child: Icon(icon, size: 32, color: primaryBackgroundColor),
              ),
            ),
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  SelectableText(title, style: Theme.of(context).textTheme.headline2),
                  SelectableText(subtitle ?? "", style: Theme.of(context).textTheme.bodyText2),
                ],
              ),
            ),
            ...(children ?? []),
          ],
        ),
      ),
    );
  }
}
