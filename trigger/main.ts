import { logger, schedules, wait, tags } from "@trigger.dev/sdk/v3";

export const firstScheduledTask = schedules.task({
  id: "first-scheduled-task",
  machine: {
    preset: "micro"
  },
  cron: {
    timezone: "Europe/Amsterdam",
    pattern: "0 */4 * * *"
  },
  retry: {
    maxAttempts: 3,
    factor: 1.9,
    minTimeoutInMs: 500,
    maxTimeoutInMs: 30000,
    randomize: false,
  },
  maxDuration: 90, // Stop axter 90s
  run: async (payload, { ctx }) => {

    await tags.add("")
  },
}); 