# AnalyticsGen

Google Sheet Analytics Generator Source Code supported Flutter, Android, and iOS

## Download

### Homebrew

```shell
$ brew update
$ brew tap prongbang/homebrew-formulae
$ brew install analyticsgen
```

### Binary

- [macOS](https://github.com/prongbang/analyticsgen/blob/master/bin/macos/analyticsgen?raw=true)
- [Linux](https://github.com/prongbang/analyticsgen/blob/master/bin/linux/analyticsgen?raw=true)
- [Mindows](https://github.com/prongbang/analyticsgen/blob/master/bin/windows/analyticsgen.exe?raw=true)


### Golang

```shell
go install github.com/prongbang/analyticsgen
```

## Requirement

- Create a Google Sheet and share with `Anyone on the Internet with this link can view`

![Sheet](/docs/screenshot.png)

- Extract from the link the `DocumentId` and `SheetId` values

```
https://docs.google.com/spreadsheets/d/<DocumentId>/edit#gid=<SheetId>
```

- Example

```html
https://docs.google.com/spreadsheets/d/1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU/edit#gid=0
```

## Usage

### Flutter

- Gen all

```shell
$ analyticsgen -platform flutter -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export -package firebasex/analytics
```

- Gen by asset

```shell
$ analyticsgen -platform flutter -asset key -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export -package firebasex/analytics
$ analyticsgen -platform flutter -asset code -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export -package firebasex/analytics
```

- Add `analytics_utility.dart` file into project

```dart
import 'package:firebase_analytics/firebase_analytics.dart';

abstract class AnalyticsUtility {
  Future<void> logAppOpen();

  Future<void> logEvent(String key, Map<String, dynamic> parameters);

  Future<void> setUserProperty(String key, String value);

  Future<void> logScreen(String screenName, {String screenClassOverride});
}

class FirebaseAnalyticsUtility implements AnalyticsUtility {
  final FirebaseAnalytics _firebaseAnalytics;

  FirebaseAnalyticsUtility(this._firebaseAnalytics);

  @override
  Future<void> logAppOpen() => _firebaseAnalytics.logAppOpen();

  @override
  Future<void> logEvent(String key, Map<String, dynamic> parameters) =>
      _firebaseAnalytics.logEvent(name: key, parameters: parameters);

  @override
  Future<void> setUserProperty(String key, String value) =>
      _firebaseAnalytics.setUserProperty(name: key, value: value);

  @override
  Future<void> logScreen(String screenName,
          {String screenClassOverride = 'Flutter'}) =>
      _firebaseAnalytics.setCurrentScreen(
        screenName: screenName,
        screenClassOverride: screenClassOverride,
      );
}
```

### iOS

- Gen all

```shell
$ analyticsgen -platform ios -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export
```

- Gen by asset

```shell
$ analyticsgen -platform ios -asset key -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export
$ analyticsgen -platform ios -asset code -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export
```

### Android

- Gen all

```shell
$ analyticsgen -platform android -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export
```

- Gen by asset

```shell
$ analyticsgen -platform android -asset key -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export
$ analyticsgen -platform android -asset code -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export
```
