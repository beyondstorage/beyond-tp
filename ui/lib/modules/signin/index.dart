import 'package:get/get.dart';
import 'package:flutter/material.dart';

import 'form.dart';
import 'controller.dart';

class SignIn extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final SigninController c = Get.put(SigninController());
    c.initUser();

    return Scaffold(
      backgroundColor: Theme.of(context).scaffoldBackgroundColor,
      body: Center(
        child: SizedBox(
          width: 400,
          child: Card(
            child: Padding(
              child: SignInForm(),
              padding: EdgeInsets.only(left: 20, right: 20, top: 40, bottom: 20),
            )
          ),
        ),
      ),
    );
  }
}
