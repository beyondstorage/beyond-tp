import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../layout/index.dart';

class PageNotFound extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Layout(
      child: Center(
        child: SizedBox(
          height: 300.0,
          child: Column(
            children: [
              Text(
                "404",
                style: Theme.of(context).textTheme.headline1,
              ),
              Text(
                "Page not found.".tr,
                style: Theme.of(context).textTheme.headline5,
              )
            ],
          ),
        ),
      )
    );
  }
}