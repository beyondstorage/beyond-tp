import 'package:flutter/material.dart';

import '../layout/index.dart';
import '../dashboard/index.dart';

class Home extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Layout(child: Dashboard());
  }
}