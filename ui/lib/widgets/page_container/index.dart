import 'package:flutter/material.dart';

class PageContainer extends StatelessWidget {
  final List<Widget> children;

  PageContainer({ this.children });

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Container(
          width: double.infinity,
          padding: EdgeInsets.only(left: 20, right: 20, bottom: 20),
          decoration: new BoxDecoration(
            color: Colors.white,
            borderRadius: BorderRadius.all(Radius.circular(2.0)),
            boxShadow: [
              BoxShadow(
                offset: Offset(0, 1),
                color: Color.fromRGBO(26, 30, 34, 0.08),
                blurRadius: 3.0,
              )
            ],
          ),
          child: Column(children: children),
        ),
      ],
    );
  }
}
