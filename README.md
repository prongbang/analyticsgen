# AnalyticsGen

Google Sheet Analytics Generator Source Code supported Flutter, Android, and iOS

## Setup

- Coming...

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

- Flutter

```shell
$ analyticsgen -platform flutter -asset key -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export
$ analyticsgen -platform flutter -asset code -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export
```

- iOS

```shell
$ analyticsgen -platform ios -asset key -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export
$ analyticsgen -platform ios -asset code -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export
```

- Android

```shell
$ analyticsgen -platform android -asset key -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export
$ analyticsgen -platform android -asset code -sheet 0 -document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU -target ./export
```
