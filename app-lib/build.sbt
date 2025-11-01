ThisBuild / scalaVersion := "3.3.6"

lazy val orgchartsdotorgAppLib = project
  .in(file("."))
  .enablePlugins(ScalaJSPlugin)
  .settings(
    libraryDependencies ++= Seq(
      "org.typelevel" %%% "cats-effect" % "3.6.3",
      "org.http4s" %%% "http4s-ember-client" % "0.23.32",
      "org.http4s" %%% "http4s-circe" % "0.23.32"
    ),
    scalaJSLinkerConfig ~= { _.withModuleKind(ModuleKind.ESModule) }
  )
