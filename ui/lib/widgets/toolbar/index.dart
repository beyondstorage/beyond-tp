import 'package:flutter/material.dart';

class PageToolbar extends StatelessWidget {
  final String title;
  final List<Widget> children;

  PageToolbar({ required this.title, required this.children });

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 56.0,
      alignment: Alignment.center,
      margin: EdgeInsets.only(bottom: 16.0),
      child: Row(
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          Expanded(
            child: SelectableText(
              title, style: Theme.of(context).textTheme.headline6),
          ),
          ...children,
        ]
      ),
      decoration: new BoxDecoration(
        border: Border(
          bottom: BorderSide(
            style: BorderStyle.solid,
            color: Color.fromRGBO(228, 235, 241, 1),
          )
        ),
      ),
    );
  }
}