/*
 * Copyright 2018-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package integration

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudfoundry/dagger"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
)

var ExpectedExtensions = [...]string{
	"bz2",
	"curl",
	"dba",
	"exif",
	"fileinfo",
	"ftp",
	"gd",
	"gettext",
	"gmp",
	"imap",
	"ldap",
	"mbstring",
	"mysqli",
	"opcache",
	"openssl",
	"pcntl",
	"pdo",
	"pdo_mysql",
	"pdo_pgsql",
	"pdo_sqlite",
	"pgsql",
	"pspell",
	"shmop",
	"snmp",
	"soap",
	"sockets",
	"sysvmsg",
	"sysvsem",
	"sysvshm",
	"xsl",
	"zip",
	"zlib",
	"apcu",
	"cassandra",
	"geoip",
	"igbinary",
	"gnupg",
	"imagick",
	"lzf",
	"mailparse",
	"mongodb",
	"msgpack",
	"oauth",
	"odbc",
	"pdo_odbc",
	"pdo_sqlsrv",
	"rdkafka",
	"redis",
	"sqlsrv",
	"stomp",
	"xdebug",
	"yaf",
	"yaml",
	"memcached",
	"sodium",
	"tidy",
	"enchant",
	"interbase",
	"pdo_firebird",
	"readline",
	"wddx",
	"xmlrpc",
	"recode",
	"amqp",
	"lua",
	"maxminddb",
	"phalcon",
	"phpiredis",
	"protobuf",
	"tideways",
	"tideways_xhprof",
	"ioncube"}


	func TestDeployAPHPAppWithAllExtensions(t *testing.T) {
	spec.Run(t, "Deploy a PHP app with all extensions", testDeployAPHPAppWithAllExtensions, spec.Report(report.Terminal{}))
}

func testDeployAPHPAppWithAllExtensions(t *testing.T, when spec.G, it spec.S) {
	var app *dagger.App

	it.Before(func() {
		var err error

		RegisterTestingT(t)

		app, err = PreparePhpApp("php_modules")
		Expect(err).ToNot(HaveOccurred())
	})


	when("deploying a basic PHP app", func() {
		it("loads key bundled extensions", func() {

			app.SetHealthCheck("true", "3s", "1s")
			err := app.Start()


			if err != nil {
				_, err = fmt.Fprintf(os.Stderr, "App failed to start: %v\n", err)
				containerID, imageName, volumeIDs, err := app.Info()
				Expect(err).NotTo(HaveOccurred())
				fmt.Printf("ContainerID: %s\nImage Name: %s\nAll leftover cached volumes: %v\n", containerID, imageName, volumeIDs)

				containerLogs, err := app.Logs()
				Expect(err).NotTo(HaveOccurred())
				fmt.Printf("Container Logs:\n %s\n", containerLogs)
				t.FailNow()
			}

			output, err :=app.Logs()


			for _, extension := range ExpectedExtensions {
				Expect(output).To(ContainSubstring(extension))
			}
		})

		it.After(func() {
			Expect(app.Destroy()).To(Succeed()) //Only destroy app if the test passed to leave artifacts to debug
		})
	})
}