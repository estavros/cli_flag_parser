import java.util.HashMap;
import java.util.Map;

public class CLIParser {

    private Map<String, String> flags = new HashMap<>();

    public CLIParser(String[] args) {
        parse(args);
    }

    private void parse(String[] args) {
        for (int i = 0; i < args.length; i++) {
            String arg = args[i];

            if (arg.startsWith("--")) { // long flag --verbose
                String key = arg.substring(2);
                if (i + 1 < args.length && !args[i + 1].startsWith("-")) {
                    flags.put(key, args[i + 1]);
                    i++; // skip next argument since it is value
                } else {
                    flags.put(key, "true"); // boolean flag
                }
            } else if (arg.startsWith("-")) { // short flag -v
                String key = arg.substring(1);
                if (i + 1 < args.length && !args[i + 1].startsWith("-")) {
                    flags.put(key, args[i + 1]);
                    i++;
                } else {
                    flags.put(key, "true");
                }
            }
        }
    }

    public boolean hasFlag(String flag) {
        return flags.containsKey(flag);
    }

    public String getFlagValue(String flag) {
        return flags.get(flag);
    }

    // Demo
    public static void main(String[] args) {
        CLIParser parser = new CLIParser(args);

        if (parser.hasFlag("verbose") || parser.hasFlag("v")) {
            System.out.println("Verbose mode is ON");
        }

        if (parser.hasFlag("file") || parser.hasFlag("f")) {
            System.out.println("File: " + parser.getFlagValue("file"));
        }

        if (parser.hasFlag("n")) {
            System.out.println("Number: " + parser.getFlagValue("n"));
        }
    }
}
