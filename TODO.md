# Fitness Backend: Prioritized TODOs

## Highest Priority

## High Priority

## Medium Priority

9. Remove redundant params/logic
   - Drop unused `routine_id` param in `/workouts/list` or apply filter.
   - Clean commented/dead code (e.g., commented `userID` filter in routine read).

10. Add input validation for DTOs
   - Non-empty names, positive `target_sets`, reasonable weight bounds.

11. Logging cleanup
   - Fix typos; add contextual fields; consider structured logs.

12. Consistency pass
   - Convert remaining `http.Error` to JSON helper.
   - Normalize headers and response patterns.