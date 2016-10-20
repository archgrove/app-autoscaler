package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
	"io/ioutil"
	"os"
)

var _ = Describe("Pruner", func() {
	var runner *PrunerRunner

	BeforeEach(func() {
		initConfig()
		runner = NewPrunerRunner()
	})

	JustBeforeEach(func() {
		runner.Start()
	})

	AfterEach(func() {
		runner.KillWithFire()
	})

	Context("when pruner started", func() {
		It("should start metricsdbpruner", func() {

			//Metric Pruner
			Eventually(runner.Session).Should(Say("metricsdbpruner.started"))

			// Pruner
			Consistently(runner.Session).ShouldNot(Exit())

		})

		It("should start appmetricsdbpruner", func() {

			//App Metric Pruner
			Eventually(runner.Session).Should(Say("appmetricsdbpruner.started"))

			// Pruner
			Consistently(runner.Session).ShouldNot(Exit())

		})

	})

	Context("with a missing config file", func() {
		BeforeEach(func() {
			runner.configPath = "bogus"
		})

		It("fails with an error", func() {
			Eventually(runner.Session).Should(Exit(1))
			Expect(runner.Session.Buffer()).To(Say("failed to open config file"))
		})
	})

	Context("with an invalid config file", func() {
		BeforeEach(func() {
			badfile, err := ioutil.TempFile("", "bad-pr-config")
			Expect(err).NotTo(HaveOccurred())
			runner.configPath = badfile.Name()
			ioutil.WriteFile(runner.configPath, []byte("bogus"), os.ModePerm)
		})

		AfterEach(func() {
			os.Remove(runner.configPath)
		})

		It("fails with an error", func() {
			Eventually(runner.Session).Should(Exit(1))
			Expect(runner.Session.Buffer()).To(Say("failed to read config file"))
		})
	})

	Context("with missing/invalid configuration", func() {
		BeforeEach(func() {

			cfg.MetricsDb.CutoffDays = -1

			cfg := writeConfig(&cfg)
			runner.configPath = cfg.Name()
		})

		AfterEach(func() {
			os.Remove(runner.configPath)
		})

		It("should fail validation", func() {
			Eventually(runner.Session).Should(Exit(1))
			Expect(runner.Session.Buffer()).To(Say("failed to validate configuration"))
		})
	})

	Context("when connection to metrics db fails", func() {
		BeforeEach(func() {

			//invalid url
			cfg.MetricsDb.DbUrl = "postgres://not-exist-user:not-exist-password@localhost/autoscaler?sslmode=disable"

			cfg := writeConfig(&cfg)
			runner.configPath = cfg.Name()

		})

		AfterEach(func() {
			os.Remove(runner.configPath)
		})

		It("should error", func() {
			Eventually(runner.Session).Should(Exit(1))
			Expect(runner.Session.Buffer()).To(Say("failed to connect metrics db"))
		})

	})

	Context("when connection to app metrics db fails", func() {
		BeforeEach(func() {

			//invalid url
			cfg.AppMetricsDb.DbUrl = "postgres://not-exist-user:not-exist-password@localhost/autoscaler?sslmode=disable"

			cfg := writeConfig(&cfg)
			runner.configPath = cfg.Name()

		})

		AfterEach(func() {
			os.Remove(runner.configPath)
		})

		It("should error", func() {
			Eventually(runner.Session).Should(Exit(1))
			Expect(runner.Session.Buffer()).To(Say("failed to connect app metrics db"))
		})

	})

	Context("when an interrupt is sent", func() {
		It("should stop", func() {
			runner.Interrupt()
			Eventually(runner.Session, 5).Should(Exit(130))
		})
	})
})