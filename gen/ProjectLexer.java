// Generated from /home/jdvelasquezo/GolandProjects/OLC2_Project1/server/ProjectLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ProjectLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PRINTLN=1, PRINT=2, DECLARAR=3, MUT=4, RSTRING=5, RINTEGER=6, RREAL=7, 
		RBOOLEAN=8, RCHAR=9, RSTR=10, RAS=11, RIF=12, RELSE=13, POWI=14, POWF=15, 
		TOSTRING=16, TOOWNED=17, MULTILINE=18, INLINE=19, INTEGER=20, FLOAT=21, 
		CHAR=22, STRING=23, BOOLEAN=24, ID=25, EQUAL=26, DOT=27, SEMICOLON=28, 
		COMMA=29, COLON=30, ADMIRATION=31, REFERENCE=32, HERITAGE=33, LEFT_PAR=34, 
		RIGHT_PAR=35, LEFT_KEY=36, RIGHT_KEY=37, LEFT_BRACKET=38, RIGHT_BRACKET=39, 
		EQUEAL_EQUAL=40, NOTEQUAL=41, GREATER_THAN=42, LESS_THAN=43, GREATER_EQUALTHAN=44, 
		LESS_EQUEALTHAN=45, MUL=46, DIV=47, MOD=48, ADD=49, SUB=50, AND=51, OR=52, 
		WHITESPACE=53;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PRINTLN", "PRINT", "DECLARAR", "MUT", "RSTRING", "RINTEGER", "RREAL", 
			"RBOOLEAN", "RCHAR", "RSTR", "RAS", "RIF", "RELSE", "POWI", "POWF", "TOSTRING", 
			"TOOWNED", "MULTILINE", "INLINE", "INTEGER", "FLOAT", "CHAR", "STRING", 
			"BOOLEAN", "ID", "EQUAL", "DOT", "SEMICOLON", "COMMA", "COLON", "ADMIRATION", 
			"REFERENCE", "HERITAGE", "LEFT_PAR", "RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", 
			"LEFT_BRACKET", "RIGHT_BRACKET", "EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN", 
			"LESS_THAN", "GREATER_EQUALTHAN", "LESS_EQUEALTHAN", "MUL", "DIV", "MOD", 
			"ADD", "SUB", "AND", "OR", "WHITESPACE", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'println'", "'print'", "'let'", "'mut'", "'String'", "'i64'", 
			"'f64'", "'bool'", "'char'", "'&str'", "'as'", "'If'", "'entonces'", 
			"'pow'", "'powf'", "'to_string'", "'to_owned'", null, null, null, null, 
			null, null, null, null, "'='", "'.'", "';'", "','", "':'", "'!'", "'&'", 
			"'::'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'=='", "'!='", "'>'", 
			"'<'", "'>='", "'<='", "'*'", "'/'", "'%'", "'+'", "'-'", "'&&'", "'||'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PRINTLN", "PRINT", "DECLARAR", "MUT", "RSTRING", "RINTEGER", "RREAL", 
			"RBOOLEAN", "RCHAR", "RSTR", "RAS", "RIF", "RELSE", "POWI", "POWF", "TOSTRING", 
			"TOOWNED", "MULTILINE", "INLINE", "INTEGER", "FLOAT", "CHAR", "STRING", 
			"BOOLEAN", "ID", "EQUAL", "DOT", "SEMICOLON", "COMMA", "COLON", "ADMIRATION", 
			"REFERENCE", "HERITAGE", "LEFT_PAR", "RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", 
			"LEFT_BRACKET", "RIGHT_BRACKET", "EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN", 
			"LESS_THAN", "GREATER_EQUALTHAN", "LESS_EQUEALTHAN", "MUL", "DIV", "MOD", 
			"ADD", "SUB", "AND", "OR", "WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public ProjectLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "ProjectLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\67\u015b\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64"+
		"\t\64\4\65\t\65\4\66\t\66\4\67\t\67\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3"+
		"\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\r\3\r\3\r\3\16\3"+
		"\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\20\3\20\3"+
		"\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3"+
		"\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\6\23\u00d3"+
		"\n\23\r\23\16\23\u00d4\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\6"+
		"\24\u00e0\n\24\r\24\16\24\u00e1\3\24\3\24\3\25\6\25\u00e7\n\25\r\25\16"+
		"\25\u00e8\3\26\6\26\u00ec\n\26\r\26\16\26\u00ed\3\26\3\26\6\26\u00f2\n"+
		"\26\r\26\16\26\u00f3\3\27\3\27\3\27\3\27\3\30\3\30\7\30\u00fc\n\30\f\30"+
		"\16\30\u00ff\13\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3"+
		"\31\5\31\u010c\n\31\3\32\3\32\7\32\u0110\n\32\f\32\16\32\u0113\13\32\3"+
		"\33\3\33\3\34\3\34\3\35\3\35\3\36\3\36\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3"+
		"\"\3#\3#\3$\3$\3%\3%\3&\3&\3\'\3\'\3(\3(\3)\3)\3)\3*\3*\3*\3+\3+\3,\3"+
		",\3-\3-\3-\3.\3.\3.\3/\3/\3\60\3\60\3\61\3\61\3\62\3\62\3\63\3\63\3\64"+
		"\3\64\3\64\3\65\3\65\3\65\3\66\6\66\u0153\n\66\r\66\16\66\u0154\3\66\3"+
		"\66\3\67\3\67\3\67\2\28\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f"+
		"\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63"+
		"\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62"+
		"c\63e\64g\65i\66k\67m\2\3\2\n\3\2\61\61\3\2\f\f\3\2\62;\3\2$$\7\2C\\a"+
		"ac|\u00d3\u00d3\u00f3\u00f3\7\2\62;C\\c|\u00d3\u00d3\u00f3\u00f3\5\2\13"+
		"\f\17\17\"\"\t\2\"#%%--/\60<<BB]_\2\u0162\2\3\3\2\2\2\2\5\3\2\2\2\2\7"+
		"\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2"+
		"\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2"+
		"\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2"+
		"\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2"+
		"\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2"+
		"\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M"+
		"\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2"+
		"\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2"+
		"\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\3o\3\2\2\2\5w\3\2\2\2\7}\3\2\2\2\t\u0081"+
		"\3\2\2\2\13\u0085\3\2\2\2\r\u008c\3\2\2\2\17\u0090\3\2\2\2\21\u0094\3"+
		"\2\2\2\23\u0099\3\2\2\2\25\u009e\3\2\2\2\27\u00a3\3\2\2\2\31\u00a6\3\2"+
		"\2\2\33\u00a9\3\2\2\2\35\u00b2\3\2\2\2\37\u00b6\3\2\2\2!\u00bb\3\2\2\2"+
		"#\u00c5\3\2\2\2%\u00ce\3\2\2\2\'\u00db\3\2\2\2)\u00e6\3\2\2\2+\u00eb\3"+
		"\2\2\2-\u00f5\3\2\2\2/\u00f9\3\2\2\2\61\u010b\3\2\2\2\63\u010d\3\2\2\2"+
		"\65\u0114\3\2\2\2\67\u0116\3\2\2\29\u0118\3\2\2\2;\u011a\3\2\2\2=\u011c"+
		"\3\2\2\2?\u011e\3\2\2\2A\u0120\3\2\2\2C\u0122\3\2\2\2E\u0125\3\2\2\2G"+
		"\u0127\3\2\2\2I\u0129\3\2\2\2K\u012b\3\2\2\2M\u012d\3\2\2\2O\u012f\3\2"+
		"\2\2Q\u0131\3\2\2\2S\u0134\3\2\2\2U\u0137\3\2\2\2W\u0139\3\2\2\2Y\u013b"+
		"\3\2\2\2[\u013e\3\2\2\2]\u0141\3\2\2\2_\u0143\3\2\2\2a\u0145\3\2\2\2c"+
		"\u0147\3\2\2\2e\u0149\3\2\2\2g\u014b\3\2\2\2i\u014e\3\2\2\2k\u0152\3\2"+
		"\2\2m\u0158\3\2\2\2op\7r\2\2pq\7t\2\2qr\7k\2\2rs\7p\2\2st\7v\2\2tu\7n"+
		"\2\2uv\7p\2\2v\4\3\2\2\2wx\7r\2\2xy\7t\2\2yz\7k\2\2z{\7p\2\2{|\7v\2\2"+
		"|\6\3\2\2\2}~\7n\2\2~\177\7g\2\2\177\u0080\7v\2\2\u0080\b\3\2\2\2\u0081"+
		"\u0082\7o\2\2\u0082\u0083\7w\2\2\u0083\u0084\7v\2\2\u0084\n\3\2\2\2\u0085"+
		"\u0086\7U\2\2\u0086\u0087\7v\2\2\u0087\u0088\7t\2\2\u0088\u0089\7k\2\2"+
		"\u0089\u008a\7p\2\2\u008a\u008b\7i\2\2\u008b\f\3\2\2\2\u008c\u008d\7k"+
		"\2\2\u008d\u008e\78\2\2\u008e\u008f\7\66\2\2\u008f\16\3\2\2\2\u0090\u0091"+
		"\7h\2\2\u0091\u0092\78\2\2\u0092\u0093\7\66\2\2\u0093\20\3\2\2\2\u0094"+
		"\u0095\7d\2\2\u0095\u0096\7q\2\2\u0096\u0097\7q\2\2\u0097\u0098\7n\2\2"+
		"\u0098\22\3\2\2\2\u0099\u009a\7e\2\2\u009a\u009b\7j\2\2\u009b\u009c\7"+
		"c\2\2\u009c\u009d\7t\2\2\u009d\24\3\2\2\2\u009e\u009f\7(\2\2\u009f\u00a0"+
		"\7u\2\2\u00a0\u00a1\7v\2\2\u00a1\u00a2\7t\2\2\u00a2\26\3\2\2\2\u00a3\u00a4"+
		"\7c\2\2\u00a4\u00a5\7u\2\2\u00a5\30\3\2\2\2\u00a6\u00a7\7K\2\2\u00a7\u00a8"+
		"\7h\2\2\u00a8\32\3\2\2\2\u00a9\u00aa\7g\2\2\u00aa\u00ab\7p\2\2\u00ab\u00ac"+
		"\7v\2\2\u00ac\u00ad\7q\2\2\u00ad\u00ae\7p\2\2\u00ae\u00af\7e\2\2\u00af"+
		"\u00b0\7g\2\2\u00b0\u00b1\7u\2\2\u00b1\34\3\2\2\2\u00b2\u00b3\7r\2\2\u00b3"+
		"\u00b4\7q\2\2\u00b4\u00b5\7y\2\2\u00b5\36\3\2\2\2\u00b6\u00b7\7r\2\2\u00b7"+
		"\u00b8\7q\2\2\u00b8\u00b9\7y\2\2\u00b9\u00ba\7h\2\2\u00ba \3\2\2\2\u00bb"+
		"\u00bc\7v\2\2\u00bc\u00bd\7q\2\2\u00bd\u00be\7a\2\2\u00be\u00bf\7u\2\2"+
		"\u00bf\u00c0\7v\2\2\u00c0\u00c1\7t\2\2\u00c1\u00c2\7k\2\2\u00c2\u00c3"+
		"\7p\2\2\u00c3\u00c4\7i\2\2\u00c4\"\3\2\2\2\u00c5\u00c6\7v\2\2\u00c6\u00c7"+
		"\7q\2\2\u00c7\u00c8\7a\2\2\u00c8\u00c9\7q\2\2\u00c9\u00ca\7y\2\2\u00ca"+
		"\u00cb\7p\2\2\u00cb\u00cc\7g\2\2\u00cc\u00cd\7f\2\2\u00cd$\3\2\2\2\u00ce"+
		"\u00cf\7\61\2\2\u00cf\u00d0\7,\2\2\u00d0\u00d2\3\2\2\2\u00d1\u00d3\n\2"+
		"\2\2\u00d2\u00d1\3\2\2\2\u00d3\u00d4\3\2\2\2\u00d4\u00d2\3\2\2\2\u00d4"+
		"\u00d5\3\2\2\2\u00d5\u00d6\3\2\2\2\u00d6\u00d7\7,\2\2\u00d7\u00d8\7\61"+
		"\2\2\u00d8\u00d9\3\2\2\2\u00d9\u00da\b\23\2\2\u00da&\3\2\2\2\u00db\u00dc"+
		"\7\61\2\2\u00dc\u00dd\7\61\2\2\u00dd\u00df\3\2\2\2\u00de\u00e0\n\3\2\2"+
		"\u00df\u00de\3\2\2\2\u00e0\u00e1\3\2\2\2\u00e1\u00df\3\2\2\2\u00e1\u00e2"+
		"\3\2\2\2\u00e2\u00e3\3\2\2\2\u00e3\u00e4\b\24\2\2\u00e4(\3\2\2\2\u00e5"+
		"\u00e7\t\4\2\2\u00e6\u00e5\3\2\2\2\u00e7\u00e8\3\2\2\2\u00e8\u00e6\3\2"+
		"\2\2\u00e8\u00e9\3\2\2\2\u00e9*\3\2\2\2\u00ea\u00ec\t\4\2\2\u00eb\u00ea"+
		"\3\2\2\2\u00ec\u00ed\3\2\2\2\u00ed\u00eb\3\2\2\2\u00ed\u00ee\3\2\2\2\u00ee"+
		"\u00ef\3\2\2\2\u00ef\u00f1\7\60\2\2\u00f0\u00f2\t\4\2\2\u00f1\u00f0\3"+
		"\2\2\2\u00f2\u00f3\3\2\2\2\u00f3\u00f1\3\2\2\2\u00f3\u00f4\3\2\2\2\u00f4"+
		",\3\2\2\2\u00f5\u00f6\7)\2\2\u00f6\u00f7\n\5\2\2\u00f7\u00f8\7)\2\2\u00f8"+
		".\3\2\2\2\u00f9\u00fd\7$\2\2\u00fa\u00fc\n\5\2\2\u00fb\u00fa\3\2\2\2\u00fc"+
		"\u00ff\3\2\2\2\u00fd\u00fb\3\2\2\2\u00fd\u00fe\3\2\2\2\u00fe\u0100\3\2"+
		"\2\2\u00ff\u00fd\3\2\2\2\u0100\u0101\7$\2\2\u0101\60\3\2\2\2\u0102\u0103"+
		"\7v\2\2\u0103\u0104\7t\2\2\u0104\u0105\7w\2\2\u0105\u010c\7g\2\2\u0106"+
		"\u0107\7h\2\2\u0107\u0108\7c\2\2\u0108\u0109\7n\2\2\u0109\u010a\7u\2\2"+
		"\u010a\u010c\7g\2\2\u010b\u0102\3\2\2\2\u010b\u0106\3\2\2\2\u010c\62\3"+
		"\2\2\2\u010d\u0111\t\6\2\2\u010e\u0110\t\7\2\2\u010f\u010e\3\2\2\2\u0110"+
		"\u0113\3\2\2\2\u0111\u010f\3\2\2\2\u0111\u0112\3\2\2\2\u0112\64\3\2\2"+
		"\2\u0113\u0111\3\2\2\2\u0114\u0115\7?\2\2\u0115\66\3\2\2\2\u0116\u0117"+
		"\7\60\2\2\u01178\3\2\2\2\u0118\u0119\7=\2\2\u0119:\3\2\2\2\u011a\u011b"+
		"\7.\2\2\u011b<\3\2\2\2\u011c\u011d\7<\2\2\u011d>\3\2\2\2\u011e\u011f\7"+
		"#\2\2\u011f@\3\2\2\2\u0120\u0121\7(\2\2\u0121B\3\2\2\2\u0122\u0123\7<"+
		"\2\2\u0123\u0124\7<\2\2\u0124D\3\2\2\2\u0125\u0126\7*\2\2\u0126F\3\2\2"+
		"\2\u0127\u0128\7+\2\2\u0128H\3\2\2\2\u0129\u012a\7}\2\2\u012aJ\3\2\2\2"+
		"\u012b\u012c\7\177\2\2\u012cL\3\2\2\2\u012d\u012e\7]\2\2\u012eN\3\2\2"+
		"\2\u012f\u0130\7_\2\2\u0130P\3\2\2\2\u0131\u0132\7?\2\2\u0132\u0133\7"+
		"?\2\2\u0133R\3\2\2\2\u0134\u0135\7#\2\2\u0135\u0136\7?\2\2\u0136T\3\2"+
		"\2\2\u0137\u0138\7@\2\2\u0138V\3\2\2\2\u0139\u013a\7>\2\2\u013aX\3\2\2"+
		"\2\u013b\u013c\7@\2\2\u013c\u013d\7?\2\2\u013dZ\3\2\2\2\u013e\u013f\7"+
		">\2\2\u013f\u0140\7?\2\2\u0140\\\3\2\2\2\u0141\u0142\7,\2\2\u0142^\3\2"+
		"\2\2\u0143\u0144\7\61\2\2\u0144`\3\2\2\2\u0145\u0146\7\'\2\2\u0146b\3"+
		"\2\2\2\u0147\u0148\7-\2\2\u0148d\3\2\2\2\u0149\u014a\7/\2\2\u014af\3\2"+
		"\2\2\u014b\u014c\7(\2\2\u014c\u014d\7(\2\2\u014dh\3\2\2\2\u014e\u014f"+
		"\7~\2\2\u014f\u0150\7~\2\2\u0150j\3\2\2\2\u0151\u0153\t\b\2\2\u0152\u0151"+
		"\3\2\2\2\u0153\u0154\3\2\2\2\u0154\u0152\3\2\2\2\u0154\u0155\3\2\2\2\u0155"+
		"\u0156\3\2\2\2\u0156\u0157\b\66\2\2\u0157l\3\2\2\2\u0158\u0159\7^\2\2"+
		"\u0159\u015a\t\t\2\2\u015an\3\2\2\2\f\2\u00d4\u00e1\u00e8\u00ed\u00f3"+
		"\u00fd\u010b\u0111\u0154\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}